package initialize

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

func taskinit()  {
	//test taak
	logs.Info("task start")
	//tk := toolbox.NewTask("test1", "0/1 * * * * *", test1)
	//tk2 := toolbox.NewTask("test2", "0/2 * * * * *", test2)
	//toolbox.AddTask("tk",tk)
	//toolbox.AddTask("tk2",tk2)
	//toolbox.StartTask()
}

func test1() error {
	// 业务
	fmt.Println("tk1")
	return nil
}

func test2() error {
	// 业务
	fmt.Println("tk2")
	return nil
}
