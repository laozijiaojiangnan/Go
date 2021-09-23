package main // 文件版本

import (
	"fmt"
	manager2 "gostudy/work/oldboy/day05/manager"
	"gostudy/work/oldboy/day05/utils"
)

func main() {
	m := manager2.NewManager()
	utils.Print(m.HelpText + "\n")

	var inp string
	for {
		utils.InpAndOut("请输入编号完成操作:", &inp)
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
			utils.Exit()
		case "7":
			utils.Print(m.HelpText + "\n")
		default:
			fmt.Print("---请输入正确的编号!---\n")
		}
	}
}
