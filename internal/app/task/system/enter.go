package system

import (
	"reflect"
)

func ScheduleTaskRun(taskMethod string) (res string, err error) {
	defer func() {
		if reflectErr := recover(); reflectErr != nil {
			err = reflectErr.(error)
		}
	}()

	var scheduleTask = ScheduleTask{}
	taskValue := reflect.ValueOf(&scheduleTask)
	infoFunc := taskValue.MethodByName(taskMethod)
	rs := infoFunc.Call([]reflect.Value{})
	for _, v := range rs {
		res += v.Interface().(string) + ","
	}
	// 去除最后一个逗号
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res, err
}
