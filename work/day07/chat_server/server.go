package main

import (
	"net"

	work "gostudy/work/day07"
)

func main() {
	s, err := net.Listen(work.NetWork, work.Address)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	manager := work.NewServerManager()
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
