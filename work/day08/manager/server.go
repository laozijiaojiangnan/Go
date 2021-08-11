package manager

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"

	"gostudy/work/day08/utils"
)

// RunServer 启动聊天服务器
func RunServer() {
	fmt.Println("服务器启动")
	s, err := net.Listen(utils.NetWork, utils.Address)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// 更新文件信息
	utils.NewFileControl().WriteFile(false)

	manager := NewServerManager()
	go manager.SendMessage()
	go manager.ListeningExit()
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
	connMap     map[string]*Connect // 存储连接
	messageChan chan []byte         // 存储消息
	exitSignal  chan os.Signal      // 存储断开信号
}

// HandleConn 处理套接字
func (s *ServerManager) HandleConn(conn net.Conn) {
	ip := conn.RemoteAddr().String()
	for {
		// 读取消息
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				// 发送注销信息
				connect, ok := s.connMap[ip]
				if ok {
					// 删除注销的conn
					s.RemoveConn(conn)
					names := s.OnlineNames()
					tx := fmt.Sprintf(
						"[系统消息]: [%s] 已注销，在线[%d]人,[%s]", connect.username, len(s.connMap), names,
					)
					s.messageChan <- []byte(tx)
				}
			}
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
			for _, connect := range s.connMap {
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

	tx := fmt.Sprintf("[%s]:%s", message.Username, message.Content)
	if message.Type == utils.Login {
		// 处理登录消息
		connect := NewConnect(conn, message.Username)
		ip := conn.RemoteAddr().String()
		s.connMap[ip] = connect

		names := s.OnlineNames()
		tx = fmt.Sprintf(
			"[系统消息]: [%s]已上线, 在线[%d]人,[%s]", message.Username, len(s.connMap), names,
		)
	}
	return []byte(tx)
}

// RemoveConn 注销 conn
func (s *ServerManager) RemoveConn(conn net.Conn) {
	ip := conn.RemoteAddr().String()
	_, ok := s.connMap[ip]
	if ok {
		delete(s.connMap, ip)
	}
}

// ListeningExit 监听退出
func (s *ServerManager) ListeningExit() {
	for {
		select {
		case <-s.exitSignal:
			// 服务端结束，把 mark 改为 true
			utils.NewFileControl().WriteFile(true)
		}
	}
}

// OnlineNames 在线人员名称
func (s *ServerManager) OnlineNames() string {
	names := ""
	for _, co := range s.connMap {
		names += co.username + ","
	}
	return strings.TrimRight(names, ",")
}

func NewServerManager() *ServerManager {
	// 初始化 signal 并绑定到 channel，一旦有信号到达，signal会发送到channel中
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	return &ServerManager{
		connMap:     make(map[string]*Connect),
		messageChan: make(chan []byte),
		exitSignal:  sig,
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
