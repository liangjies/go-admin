package utils

import (
	"io/ioutil"
	"sync"
)

type IPQuery struct {
	Data    []byte
	DataLen uint32
	IpCache sync.Map
}

// LoadData 从内存加载IP数据库
func (ipQuery *IPQuery) LoadData(database []byte) {
	ipQuery.Data = database
	ipQuery.DataLen = uint32(len(database))
}

// LoadFile 从文件加载IP数据库
func (ipQuery *IPQuery) LoadFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	ipQuery.Data = data
	ipQuery.DataLen = uint32(len(data))
}
