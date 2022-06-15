package system

import (
	"go-admin/internal/app/global"
	"go-admin/internal/app/model/common/request"
	"go-admin/internal/app/model/system"
	systemReq "go-admin/internal/app/model/system/request"
)

type SysUserLoginService struct{}

//@function: CreateSysUserLogin
//@description: 创建记录
//@param: SysUserLogin model.SysUserLogin
//@return: err error
func (sysUserLoginService *SysUserLoginService) CreateSysUserLogin(sysUserLogin system.SysUserLogin) (err error) {
	err = global.SYS_DB.Create(&sysUserLogin).Error
	return err
}

//@function: DeleteSysUserLoginByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error

func (sysUserLoginService *SysUserLoginService) DeleteSysUserLoginByIds(ids request.IdsReq) (err error) {
	err = global.SYS_DB.Delete(&[]system.SysUserLogin{}, "id in (?)", ids.Ids).Error
	return err
}

//@function: DeleteSysUserLogin
//@description: 删除操作记录
//@param: SysUserLogin model.SysUserLogin
//@return: err error

func (sysUserLoginService *SysUserLoginService) DeleteSysUserLogin(sysUserLogin system.SysUserLogin) (err error) {
	err = global.SYS_DB.Delete(&sysUserLogin).Error
	return err
}

//@function: DeleteSysUserLogin
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, SysUserLogin model.SysUserLogin

func (sysUserLoginService *SysUserLoginService) GetSysUserLogin(id uint) (err error, sysUserLogin system.SysUserLogin) {
	err = global.SYS_DB.Where("id = ?", id).First(&sysUserLogin).Error
	return
}

//@function: GetSysUserLoginInfoList
//@description: 分页获取操作记录列表
//@param: info systemReq.SysUserLoginSearch
//@return: err error, list interface{}, total int64

func (sysUserLoginService *SysUserLoginService) GetSysUserLoginInfoList(info systemReq.SysUserLoginSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.SYS_DB.Model(&system.SysUserLogin{})
	var SysUserLogins []system.SysUserLogin
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("username  LIKE ?", "%"+info.Username+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+info.Ip+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&SysUserLogins).Error
	return err, SysUserLogins, total
}
