package qqwry

import (
	"fmt"
	"go-admin/internal/app/global"
	"testing"
)

var ipQuery IPQuery

func init() {
	if err := ipQuery.LoadFile("..\\..\\..\\..\\assets\\qqwry.dat"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipQuery.DataLen)
	global.SYS_IPQuery = &ipQuery
}

func TestQQWry(t *testing.T) {
	// var ipQuery IPQuery
	// if err := ipQuery.LoadFile("..\\..\\..\\assets\\qqwry.dat"); err != nil {
	// 	fmt.Println(err)
	// }
	city, isp, err := global.SYS_IPQuery.QueryIP("192.168.1.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(city, isp)
}

// func fib() {
// 	city, isp, err := global.SYS_IPQuery.QueryIP("192.168.1.1")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	_, _ = city, isp
// }
// func BenchmarkFib(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		fib() // run fib(30) b.N times
// 	}
// }
