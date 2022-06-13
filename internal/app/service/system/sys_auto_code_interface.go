package system

import (
	"go-admin/internal/app/model/system/response"
)

type Database interface {
	GetDB() (data []response.Db, err error)
	GetTables(dbName string) (data []response.Table, err error)
	GetColumn(tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database() Database {
	return AutoCodeMysql
}
