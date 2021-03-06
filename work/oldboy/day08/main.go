package main

import (
	manager2 "gostudy/work/oldboy/day08/manager"
	utils2 "gostudy/work/oldboy/day08/utils"
)

func main() {
	/* 新增功能
	1. 增加在线人数，和对应的名字 ✅
	2. 客户端断开给所有人发送消息 ✅
	3. 启动客户端和服务端使用一个main文件 ✅
	*/

	/* 实现细节
	1. []或者{}里存了多少套接字，就有多少在线人数，循环[]或者{}就能拿到所有用户名
	2. 客户端断开，服务端会收到一个error，监听收到以后，广播消息 并去[]或者{}里删除conn
	3. 因为client和server都在一个文件里，所以需要一个条件来判断在启动的时候运行谁
	现在的做法是，把这个条件放到一个文件里，如果这个条件为true那就运行server，如果为false运行client
	并在服务器结束的时候把true改回false
	*/

	if utils2.NewFileControl().GetMark() {
		manager2.RunServer()
	} else {
		manager2.RunClient()
	}
}
