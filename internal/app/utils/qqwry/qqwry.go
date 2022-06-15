package qqwry

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net"
	"strings"
	"sync"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type IPQuery struct {
	Data    []byte
	DataLen uint32
	IpCache sync.Map
}
type cache struct {
	City string
	Isp  string
}

const (
	indexLen      = 7
	redirectMode1 = 0x01
	redirectMode2 = 0x02
)

func byte3ToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}

func gb18030Decode(src []byte) string {
	in := bytes.NewReader(src)
	out := transform.NewReader(in, simplifiedchinese.GB18030.NewDecoder())
	d, _ := ioutil.ReadAll(out)
	return string(d)
}

// LoadData 从内存加载IP数据库
func (ipQuery *IPQuery) LoadData(database []byte) {
	ipQuery.Data = database
	ipQuery.DataLen = uint32(len(database))
}

// LoadFile 从文件加载IP数据库
func (ipQuery *IPQuery) LoadFile(filepath string) (err error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	ipQuery.Data = data
	ipQuery.DataLen = uint32(len(data))
	return
}

// QueryIP 从内存或缓存查询IP
func (ipQuery *IPQuery) QueryIP(queryIp string) (city string, isp string, err error) {
	if v, ok := ipQuery.IpCache.Load(queryIp); ok {
		city = v.(cache).City
		isp = v.(cache).Isp
		return
	}
	ip := net.ParseIP(queryIp).To4()
	if ip == nil {
		err = errors.New("ip is not ipv4")
		return
	}
	ip32 := binary.BigEndian.Uint32(ip)
	posA := binary.LittleEndian.Uint32(ipQuery.Data[:4])
	posZ := binary.LittleEndian.Uint32(ipQuery.Data[4:8])
	var offset uint32 = 0
	for {
		mid := posA + (((posZ-posA)/indexLen)>>1)*indexLen
		buf := ipQuery.Data[mid : mid+indexLen]
		_ip := binary.LittleEndian.Uint32(buf[:4])
		if posZ-posA == indexLen {
			offset = byte3ToUInt32(buf[4:])
			buf = ipQuery.Data[mid+indexLen : mid+indexLen+indexLen]
			if ip32 < binary.LittleEndian.Uint32(buf[:4]) {
				break
			} else {
				offset = 0
				break
			}
		}
		if _ip > ip32 {
			posZ = mid
		} else if _ip < ip32 {
			posA = mid
		} else if _ip == ip32 {
			offset = byte3ToUInt32(buf[4:])
			break
		}
	}
	if offset <= 0 {
		err = errors.New("ip not found")
		return
	}
	posM := offset + 4
	mode := ipQuery.Data[posM]
	var ispPos uint32
	switch mode {
	case redirectMode1:
		posC := byte3ToUInt32(ipQuery.Data[posM+1 : posM+4])
		mode = ipQuery.Data[posC]
		posCA := posC
		if mode == redirectMode2 {
			posCA = byte3ToUInt32(ipQuery.Data[posC+1 : posC+4])
			posC += 4
		}
		for i := posCA; i < ipQuery.DataLen; i++ {
			if ipQuery.Data[i] == 0 {
				city = string(ipQuery.Data[posCA:i])
				break
			}
		}
		if mode != redirectMode2 {
			posC += uint32(len(city) + 1)
		}
		ispPos = posC
	case redirectMode2:
		posCA := byte3ToUInt32(ipQuery.Data[posM+1 : posM+4])
		for i := posCA; i < ipQuery.DataLen; i++ {
			if ipQuery.Data[i] == 0 {
				city = string(ipQuery.Data[posCA:i])
				break
			}
		}
		ispPos = offset + 8
	default:
		posCA := offset + 4
		for i := posCA; i < ipQuery.DataLen; i++ {
			if ipQuery.Data[i] == 0 {
				city = string(ipQuery.Data[posCA:i])
				break
			}
		}
		ispPos = offset + uint32(5+len(city))
	}
	if city != "" {
		city = strings.TrimSpace(gb18030Decode([]byte(city)))
	}
	ispMode := ipQuery.Data[ispPos]
	if ispMode == redirectMode1 || ispMode == redirectMode2 {
		ispPos = byte3ToUInt32(ipQuery.Data[ispPos+1 : ispPos+4])
	}
	if ispPos > 0 {
		for i := ispPos; i < ipQuery.DataLen; i++ {
			if ipQuery.Data[i] == 0 {
				isp = string(ipQuery.Data[ispPos:i])
				if isp != "" {
					if strings.Contains(isp, "CZ88.NET") {
						isp = ""
					} else {
						isp = strings.TrimSpace(gb18030Decode([]byte(isp)))
					}
				}
				break
			}
		}
	}
	ipQuery.IpCache.Store(queryIp, cache{City: city, Isp: isp})
	return
}
