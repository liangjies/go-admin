package system

import (
	"go-admin/internal/app/global"
)

type SysJwtBlacklist struct {
	global.SYS_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
