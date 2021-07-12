package manager

type Student struct {
	ID   string `json:"id"`   // ID
	Name string `json:"name"` // 姓名
	Score
}

type Score struct {
	Chinese int `json:"chinese"` // 语文成绩
	English int `json:"english"` // 英语成绩
}

func (s *Student) Update(name string, chinese, english int) {
	s.Name = name
	s.Chinese = chinese
	s.English = english
}
