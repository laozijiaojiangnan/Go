package manager

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	u "gostudy/work/day05/utils"
)

type Manager struct {
	Stus     map[string]*Student // 学生列表
	HelpText string              // 帮助文本
	FilePath string              // 文件路径
}

func NewManager() *Manager {
	filePath := "./students.json"
	stus := make(map[string]*Student)

	f, _ := os.Open(filePath)
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	_ = json.Unmarshal(bytes, &stus)
	return &Manager{
		stus,
		`脑残管理系统:
  1.查询学生列表
  2.添加学生
  3.更新学生
  4.删除学生
  5.保存数据
  6.直接退出
  7.显示帮助`,
		filePath,
	}

}

func (m Manager) Add() {
	u.Print("---添加学生中---\n")
	var (
		id, name string
		ch, en   int
	)

	newStudent := func(id, name string, chinese, english int) *Student {
		return &Student{
			id,
			name,
			Score{chinese, english},
		}
	}

	u.InpAndOut("请输入ID:", &id)
	u.InpAndOut("请输入姓名:", &name)
	u.InpAndOut("请输入语文成绩：", &ch)
	u.InpAndOut("请输入英语成绩：", &en)

	m.Stus[id] = newStudent(id, name, ch, en)

	u.Print("---添加成功!---\n")
}

func (m Manager) Delete() {
	u.Print("---删除学生中---\n")
	var id string
	u.InpAndOut("输入需要删除的学生ID:", &id)

	_, ok := m.Stus[id]
	if !ok {
		u.Print("未找到ID对应的学生\n")
	} else {
		delete(m.Stus, id)
		u.Print("---删除成功---\n")
	}
}

func (m Manager) Update() {
	u.Print("---更新学生中---\n")
	var (
		id, name string
		ch, en   int
	)
	u.InpAndOut("输入需要更新的学生ID:", &id)

	student, ok := m.Stus[id]
	if !ok {
		u.Print("未找到ID对应的学生\n")
	} else {
		u.InpAndOut("新名称:", &name)
		u.InpAndOut("新语文成绩:", &ch)
		u.InpAndOut("新英语成绩:", &en)
		student.Update(name, ch, en)
		u.Print("---更新成功---\n")
	}
}

func (m Manager) List() {
	if len(m.Stus) < 1 {
		u.Print("---暂时没有学生---\n")
		return
	}

	u.Print("------------------------------------------------------\n")
	for id, stu := range m.Stus {
		fmt.Printf(
			"ID:%v, 姓名:%v, 语文:%v, 数学:%v\n", id, stu.Name, stu.Chinese, stu.English,
		)
	}
	u.Print("------------------------------------------------------\n")
}

func (m Manager) Save() {
	// 打开文件，如果不存在就创建
	file, err := os.OpenFile(m.FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(m.FilePath)
	}
	defer file.Close()

	jsonData, err := json.Marshal(m.Stus)
	if err != nil {
		fmt.Printf("序列化错误:%v", err)
		return
	}

	_, err = io.WriteString(file, string(jsonData))
	if err != nil {
		fmt.Printf("写文件错误:%v", err)
		return
	}

	u.Print("---保存数据成功---\n")
}
