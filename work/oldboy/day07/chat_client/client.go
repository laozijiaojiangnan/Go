package main

import (
	"gostudy/work/oldboy/day07"
	"net"
)

func main() {
	conn, err := net.Dial(day07.NetWork, day07.Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	manager := day07.NewClientManager(conn)

	// 接受数据
	go manager.Read()
	// 发送数据
	manager.Write()
}
