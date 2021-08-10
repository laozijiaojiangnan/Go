package manager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Scan() []byte {
	fmt.Println("请输入用户名")

	reader := bufio.NewReader(os.Stdin)
	data := make([]byte, 1024)
	n, err := reader.Read(data)
	if err != nil {
		panic(err)
	}
	return data[:n]
}

func Marshal(t, text, username string) []byte {
	message := NewMessage(t, text, username)
	data, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	return data
}

func UnMarshal(data []byte) *Message {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		panic(err)
	}
	return &message
}
