package task

import beegoTask "github.com/beego/beego/v2/task"

func StartTask() {
	tk1 := CheckRelayStatusTask()
	beegoTask.AddTask("CheckNostrRelayStatusTask", tk1)

	// start tasks
	beegoTask.StartTask()
}
