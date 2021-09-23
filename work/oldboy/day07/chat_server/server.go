package main

import (
	"gostudy/work/oldboy/day07"
	"net"
)

func main() {
	s, err := net.Listen(day07.NetWork, day07.Address)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	manager := day07.NewServerManager()
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
