package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Manager 管理类
type Manager struct {
	stus     map[string]*Student // 学生列表
	helpText string              // 帮助文本
	filePath string              // 文件路径
}

// Student 学生类
type Student struct {
	ID   string `json:"id"`   // ID
	Name string `json:"name"` // 姓名
	Score
}

// Score 成绩类
type Score struct {
	Chinese int `json:"chinese"` // 语文成绩
	English int `json:"english"` // 英语成绩
}

func main() {
	m := newManager()
	_print(m.helpText + "\n")

	var inp string
	for {
		inpAndOut("请输入编号完成操作:", &inp)
		switch inp {
		case "1":
			m.list()
		case "2":
			m.add()
		case "3":
			m.update()
		case "4":
			m.delete()
		case "5":
			m.save()
		case "6":
			exit()
		case "7":
			_print(m.helpText + "\n")
		default:
			fmt.Print("---请输入正确的编号!---\n")
		}
	}
}

func newManager() *Manager {
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

//add 添加学生
func (m Manager) add() {
	_print("---添加学生中---\n")
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

	inpAndOut("请输入ID:", &id)
	inpAndOut("请输入姓名:", &name)
	inpAndOut("请输入语文成绩：", &ch)
	inpAndOut("请输入英语成绩：", &en)

	m.stus[id] = newStudent(id, name, ch, en)

	_print("---添加成功!---\n")
}

//delete 删除学生
func (m Manager) delete() {
	_print("---删除学生中---\n")
	var id string
	inpAndOut("输入需要删除的学生ID:", &id)

	_, ok := m.stus[id]
	if !ok {
		_print("未找到ID对应的学生\n")
	} else {
		delete(m.stus, id)
		_print("---删除成功---\n")
	}
}

//update 更新学生
func (m Manager) update() {
	_print("---更新学生中---\n")
	var (
		id, name string
		ch, en   int
	)
	inpAndOut("输入需要更新的学生ID:", &id)

	student, ok := m.stus[id]
	if !ok {
		_print("未找到ID对应的学生\n")
	} else {
		inpAndOut("新名称:", &name)
		inpAndOut("新语文成绩:", &ch)
		inpAndOut("新英语成绩:", &en)
		student.update(name, ch, en)
		_print("---更新成功---\n")
	}
}

//list 学生列表
func (m Manager) list() {
	if len(m.stus) < 1 {
		_print("---暂时没有学生---\n")
		return
	}

	_print("------------------------------------------------------\n")
	for id, stu := range m.stus {
		fmt.Printf(
			"ID:%v, 姓名:%v, 语文:%v, 数学:%v\n", id, stu.Name, stu.Chinese, stu.English,
		)
	}
	_print("------------------------------------------------------\n")
}

// save 保存学生信息
func (m Manager) save() {
	// 打开文件，如果不存在就创建
	file, err := os.OpenFile(m.filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(m.filePath)
	}
	defer file.Close()

	jsonData, err := json.Marshal(m.stus)
	if err != nil {
		fmt.Printf("序列化错误:%v", err)
		return
	}

	_, err = io.WriteString(file, string(jsonData))
	if err != nil {
		fmt.Printf("写文件错误:%v", err)
		return
	}

	_print("---保存数据成功---\n")
}

// update 更新学生属性
func (s *Student) update(name string, chinese, english int) {
	s.Name = name
	s.Chinese = chinese
	s.English = english
}

//exit 退出程序
func exit() {
	os.Exit(1)
}

//_print 输出函数
func _print(text string) {
	fmt.Print(text)
}

//_scanln 输入函数
func _scanln(text interface{}) {
	fmt.Scanln(text)
}

//inpAndOut 输入 + 输出函数
func inpAndOut(t string, c interface{}) {
	_print(t)
	_scanln(c)
}
