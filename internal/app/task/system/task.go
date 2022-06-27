package system

type ScheduleTask struct{}

func (scheduleTask *ScheduleTask) ScheduleTest() (res string) {
	return "ScheduleTest"
}
