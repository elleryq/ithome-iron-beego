package tasks

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	logs.Debug("Register tasks")
	periodicTask := toolbox.NewTask("periodicTask", "*/2 * * * * *", func() error {
		logs.Debug("periodicTask")
		return nil
	})
	toolbox.AddTask("periodicTask", periodicTask)
	toolbox.StartTask()
	defer toolbox.StopTask()
}
