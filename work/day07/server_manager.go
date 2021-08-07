package day07

import (
	"fmt"
	"net"
)

type ServerManager struct {
	connList    []*Connect
	messageChan chan []byte
}

func (c *ServerManager) HandleConn(conn net.Conn) {
	// 保存套接字
	remoteAddr := conn.RemoteAddr().String()
	connect := NewConnect(remoteAddr, conn)
	c.connList = append(c.connList, connect)

	// 发送默认消息
	defaultMessage := fmt.Sprintf("【系统消息】:%s 上线了, 当前一共有【%d】人在线", remoteAddr, len(c.connList))
	c.messageChan <- []byte(defaultMessage)

	for {
		// 监听客户端发送的消息
		message := make([]byte, 1024)
		n, err := conn.Read(message)
		if err != nil {
			return
		}

		// 保存消息
		msg := fmt.Sprintf("【%s】:%s", remoteAddr, string(message[:n]))
		c.messageChan <- []byte(msg)
	}
}

func (c *ServerManager) SendMessage() {
	for {
		select {
		case m := <-c.messageChan:
			for _, connect := range c.connList {
				_, err := connect.conn.Write(m)
				if err != nil {
					continue
				}
			}
		}
	}
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		connList:    []*Connect{},
		messageChan: make(chan []byte),
	}
}

type Connect struct {
	ip   string
	conn net.Conn
}

func NewConnect(ip string, conn net.Conn) *Connect {
	return &Connect{
		ip:   ip,
		conn: conn,
	}
}
