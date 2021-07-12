package main // 文件版本

import (
	"fmt"

	"gostudy/work/day05/manager"
	u "gostudy/work/day05/utils"
)

func main() {
	m := manager.NewManager()
	u.Print(m.HelpText + "\n")

	var inp string
	for {
		u.InpAndOut("请输入编号完成操作:", &inp)
		switch inp {
		case "1":
			m.List()
		case "2":
			m.Add()
		case "3":
			m.Update()
		case "4":
			m.Delete()
		case "5":
			m.Save()
		case "6":
			u.Exit()
		case "7":
			u.Print(m.HelpText + "\n")
		default:
			fmt.Print("---请输入正确的编号!---\n")
		}
	}
}
