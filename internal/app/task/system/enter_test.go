package system

import (
	"reflect"
	"testing"
)

func TestJob(t *testing.T) {
	var scheduleTask = ScheduleTask{}
	taskValue := reflect.ValueOf(&scheduleTask)
	infoFunc := taskValue.MethodByName("ScheduleTest")
	infoFunc.Call([]reflect.Value{})
	// 执行定时任务
	// ScheduleTaskRun("ScheduleTest")
}
