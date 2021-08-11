package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Control 控制运行程序的接口
type Control interface {
	GetMark() bool
}

// FileControl 使用文件控制
type FileControl struct {
	Path string
}

func (f *FileControl) GetMark() bool {
	content := f.OpenFile()

	temp := true
	if len(content) > 0 {
		var fileContent FileContent
		err := json.Unmarshal(content, &fileContent)
		if err != nil {
			panic(err)
		}
		temp = fileContent.Mark
	}
	return temp
}

func (f *FileControl) OpenFile() []byte {
	file, err := os.OpenFile(f.Path, os.O_RDWR, 0644)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(f.Path)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return data
}

func (f *FileControl) WriteFile(b bool) {
	fileContent := NewFileContent(b)
	data, err := json.Marshal(fileContent)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(f.Path, data, 0666)
	if err != nil {
		panic(err)
	}
}

func NewFileControl() *FileControl {
	return &FileControl{
		Path: "./control.json",
	}
}

type FileContent struct {
	Mark bool `json:"mark"` // 如果为true运行server，false运行client
}

func NewFileContent(mark bool) *FileContent {
	return &FileContent{
		Mark: mark,
	}
}
