package manager

import (
	"fmt"
	utils2 "gostudy/work/oldboy/day08/utils"
	"net"
)

// RunClient 运行客户端
func RunClient() {
	fmt.Println("客户端启动")
	conn, err := net.Dial(utils2.NetWork, utils2.Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Print("输入用户名:")
	username := string(Scan())
	manager := NewClientManager(conn, username)
	// 接受数据
	go manager.Read()
	// 登录
	manager.Login(username)
	// 开始聊天
	manager.Write()
}

type ClientManager struct {
	conn     net.Conn
	username string
}

func (c *ClientManager) Read() {
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

func (c *ClientManager) Write() {
	for {
		content := string(Scan())
		data := Marshal(utils2.Chat, content, c.username)

		_, err := c.conn.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

func (c *ClientManager) Login(username string) {
	data := Marshal(utils2.Login, "", username)

	_, err := c.conn.Write(data)
	if err != nil {
		panic(err)
	}
}

func NewClientManager(conn net.Conn, username string) *ClientManager {
	return &ClientManager{
		conn:     conn,
		username: username,
	}
}

type Message struct {
	Type     string `json:"type"`     // 消息类型
	Content  string `json:"content"`  // 消息内容
	Username string `json:"username"` // 用户名
}

func NewMessage(t, connect, username string) *Message {
	return &Message{
		Type:     t,
		Content:  connect,
		Username: username,
	}
}
