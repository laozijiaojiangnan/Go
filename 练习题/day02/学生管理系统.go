package main

import (
	"fmt"
	"os"
)

func main() {
	text := `
管理系统:
1.查询学生列表 ✅
2.添加学生 ✅
3.更新学生 
4.删除学生 
5.直接退出 ✅
6.显示帮助 ✅
`
	_print(text)
	var inp string
	stus := make(map[string]map[string]string)

	// 程序开始
	for {
		_print("请输入编号完成操作:")
		_scanln(&inp)
		switch inp {
		case "1": StuList(stus)
		case "2": addStu(stus)
		case "3":
		case "4":
		case "5": exit()
		case "6": _print(text)
		default:
			fmt.Print("请输入正确的编号")
		}
	}
}

//addStu 添加学生
func addStu(stus map[string]map[string]string)  {
	var id , name, age, level string
	inpAndOut("请输入ID:", &id)
	inpAndOut("请输入名称:", &name)
	inpAndOut("请输入年龄:", &age)
	inpAndOut("请输入级别:", &level)
	stus[id] = map[string]string{
		"id":id,
		"name":name,
		"age":age,
		"level":level,
	}
	_print("---添加成功!---\n")
}

//deleteStu 删除学生
func deleteStu()  {
	
}

//updateStu 更新学生
func updateStu()  {
	
}

//StuList 学生列表
func StuList(stus map[string]map[string]string)  {
	if len(stus) < 1 {
		_print("---暂时没有学生---\n")
		return
	}
	for id, item := range stus {
		fmt.Printf(
		"学生编号:%v, 学生名称:%v, 学生年龄:%v, 学生级别:%v\n", id, item["name"], item["age"], item["level"],
		)
	}
}

//exit 退出程序
func exit() {
	os.Exit(1)
}

//_print 输出函数
func _print(text string){
	fmt.Print(text)
}

//_scanln 输入函数
func _scanln(text *string)  {
	fmt.Scanln(text)
}

//inpAndOut 输入+输出函数
func inpAndOut (t string, c *string)  {
	_print(t)
	_scanln(c)
}
