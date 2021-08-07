package main

import (
	"net"

	work "gostudy/work/day07"
)

func main() {
	conn, err := net.Dial(work.NetWork, work.Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	manager := work.NewClientManager(conn)

	// 接受数据
	go manager.Read()
	// 发送数据
	manager.Write()
}
