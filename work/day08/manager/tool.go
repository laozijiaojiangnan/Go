package manager

import (
	"bufio"
	"encoding/json"
	"os"
)

func Scan() []byte {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return data
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
