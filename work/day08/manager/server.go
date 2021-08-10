package manager

import (
	"fmt"
	"net"
	"strings"

	"gostudy/work/day08/utils"
)

// RunServer 启动聊天服务器
func RunServer() {
	s, err := net.Listen(utils.NetWork, utils.Address)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// 更新文件信息
	utils.NewFileControl().WriteFile()

	manager := NewServerManager()
	go manager.SendMessage()
	for {
		var conn net.Conn
		conn, err = s.Accept()
		if err != nil {
			return
		}

		go manager.HandleConn(conn)
	}
}

type ServerManager struct {
	connList    []*Connect
	messageChan chan []byte
}

// HandleConn 处理套接字
func (s *ServerManager) HandleConn(conn net.Conn) {
	for {
		// 读取消息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			return
		}

		// 生成信息
		message := s.GenerateMessage(conn, data[:n])

		// 保存消息
		s.messageChan <- message
	}
}

// SendMessage 广播接收到的 message
func (s *ServerManager) SendMessage() {
	for {
		select {
		case m := <-s.messageChan:
			for _, connect := range s.connList {
				_, err := connect.conn.Write(m)
				if err != nil {
					continue
				}
			}
		}
	}
}

// GenerateMessage 生成消息
func (s *ServerManager) GenerateMessage(conn net.Conn, data []byte) []byte {
	message := UnMarshal(data)

	tx := fmt.Sprintf("【%s】: %s", message.Username, message.Content)
	if message.Type == utils.Login {
		// 处理登录消息
		connect := NewConnect(conn, message.Username)
		s.connList = append(s.connList, connect)

		names := ""
		for _, co := range s.connList {
			names += co.username + ","
		}
		strings.TrimRight(names, ",")

		tx = fmt.Sprintf(
			"【系统消息】%s 已上线, 在线%d人,【%s】", message.Username, len(s.connList), names,
		)
	}
	return []byte(tx)
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		connList:    []*Connect{},
		messageChan: make(chan []byte),
	}
}

type Connect struct {
	conn     net.Conn
	username string
}

func NewConnect(conn net.Conn, username string) *Connect {
	return &Connect{
		conn:     conn,
		username: username,
	}
}
