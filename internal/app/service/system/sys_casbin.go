package system

import (
	"errors"
	"fmt"
	"sync"

	"go-admin/internal/app/global"
	"go-admin/internal/app/model/system/request"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	redisadapter "github.com/casbin/redis-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

//@function: UpdateCasbin
//@description: 更新casbin权限
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	// 更新全局变量
	global.SYS_Enforcer = e
	return nil
}

//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.SYS_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		// 更新全局变量
		casbinService.CasbinInit()
	}
	return err
}

//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.SYS_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.SYS_CONFIG.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

//@function: CasbinRedis
//@description: 从Redis获取Enforcer  引入自定义规则
//@return: *casbin.Enforcer

// var (
// 	syncedEnforcer *casbin.SyncedEnforcer
// 	once           sync.Once
// )

func (casbinService *CasbinService) CasbinRedis() *casbin.SyncedEnforcer {
	once.Do(func() {
		a := redisadapter.NewAdapterWithPassword("tcp", global.SYS_CONFIG.Redis.Addr, global.SYS_CONFIG.Redis.Password)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.SYS_CONFIG.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

//@function: CasbinInitRedis
//@description: GORM到Redis  引入自定义规则
//@return: *casbin.Enforcer
func (casbinService *CasbinService) CasbinInitRedis() (err error) {
	// 从 A 加载策略到内存
	a, _ := gormadapter.NewAdapterByDB(global.SYS_DB)
	e, _ := casbin.NewEnforcer(global.SYS_CONFIG.Casbin.ModelPath, a)
	_ = e.LoadPolicy()
	//  适配器A 转换为 B
	// b := redisadapter.NewAdapterWithPassword("tcp", global.SYS_CONFIG.Redis.Addr, global.SYS_CONFIG.Redis.Password)
	fmt.Println(global.SYS_CONFIG.Redis.Addr)
	b := redisadapter.NewAdapter("tcp", global.SYS_CONFIG.Redis.Addr)
	e.SetAdapter(b)
	// // 将策略从内存保存到 B
	err = e.SavePolicy()
	return err
}

//@function: CasbinInitRedis
//@description: 初始化Casbin到全局变量
//@return: *casbin.Enforcer
func (casbinService *CasbinService) CasbinInit() {
	global.SYS_Enforcer = casbinService.Casbin()
}
