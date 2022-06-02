package task

import "github.com/astaxie/beego/toolbox"

func StartTask() {
	NewBlockTask()
	NewTransactionTask()
	NewAddressTask()
	NewBridgeTask()
	toolbox.StartTask()
}

func StopTask() {
	toolbox.StopTask()
}
