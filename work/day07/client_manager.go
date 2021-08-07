package day07

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type ClientManager struct {
	conn net.Conn
}

func (c ClientManager) Read() {
	for {
		message := make([]byte, 1024)
		n, err := c.conn.Read(message)
		if err != nil {
			return
		}
		message = message[:n]
		fmt.Println(string(message))
	}
}

func (c ClientManager) Write() {
	for {
		// 监听用户输入
		var inp string
		bufio.NewReader(os.Stdin)
		_, err := fmt.Scan(&inp)
		if err != nil {
			panic(err)
		}

		// todo 用户输入未作判断

		// 处理用户输入并发送
		_, err = c.conn.Write([]byte(inp))
		if err != nil {
			panic(err)
		}
	}
}

func NewClientManager(conn net.Conn) *ClientManager {
	return &ClientManager{
		conn: conn,
	}
}
