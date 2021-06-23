package main

import (
	"fmt"
	"os"
)

var stus = make(map[string]map[string]string)

func main() {
	text := `
管理系统:
1.查询学生列表
2.添加学生
3.更新学生
4.删除学生
5.直接退出
6.显示帮助
`
	_print(text)
	var inp string

	// 程序开始 todo 未处理输入为空的情况
	for {
		inpAndOut("请输入编号完成操作:", &inp)
		switch inp {
		case "1": StuList()
		case "2": addStu()
		case "3": updateStu()
		case "4": deleteStu()
		case "5": exit()
		case "6": _print(text)
		default:
			fmt.Print("---请输入正确的编号!---\n")
		}
	}
}

//addStu 添加学生
func addStu()  {
	_print("---添加学生中---\n")
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
	_print("---删除学生中---\n")
	var id string
	inpAndOut("输入需要删除的学生ID:", &id)

	_, ok := stus[id]
	if !ok {
		_print("未找到ID对应的学生\n")
	}else{
		delete(stus, id)
		_print("---删除成功---\n")
	}
}

//updateStu 更新学生
func updateStu()  {
	_print("---更新学生中---\n")
	var id , name, age, level string
	inpAndOut("输入需要更新的学生ID:", &id)

	stu, ok := stus[id]
	if !ok {
		_print("未找到ID对应的学生\n")
	}else{
		inpAndOut("请输入更新名称:", &name)
		inpAndOut("请输入更新年龄:", &age)
		inpAndOut("请输入更新级别:", &level)
		stu["name"] = name
		stu["age"] = age
		stu["level"] = level
		_print("---更新成功---\n")
	}
}

//StuList 学生列表
func StuList()  {
	if len(stus) < 1 {
		_print("---暂时没有学生---\n")
		return
	}
	_print("---查询学生中---\n")
	for id, item := range stus {
		fmt.Printf(
		"编号:%v, 名称:%v, 年龄:%v, 级别:%v\n", id, item["name"], item["age"], item["level"],
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

//inpAndOut 输入 + 输出函数
func inpAndOut (t string, c *string)  {
	_print(t)
	_scanln(c)
}
