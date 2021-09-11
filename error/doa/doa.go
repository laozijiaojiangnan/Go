package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	/*
	问：
		我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
		为什么，应该怎么做请写出代码？
	答：
		首先我觉得ErrNoRows更像一个标识，因为没查到数据是非常有可能发生的，在某些场合也是允许的
		所以我觉的分两种情况：
		1. 业务场景接受查询到空数据，比如一个人的购物车是可以为空的
		2. 反过来，不能接受为空数据，比如查购物车之前，我要知道当前购物车是哪个用户，这个查询就是不能为空的

		第一种情况:
			都支持为空了，这个 ErrNoRows 可以直接忽略
		第二种情况:
			我举得应该使用 Wrap 来处理错误，因为可以加额外的信息，比如什么样的数据是空，而不是只告诉ta查询为空
			这也是为了调用者考虑，如果ta想打印日志，这时候的信息是非常完整的
	*/

	ret, err := GetUsers()
	if err != nil {
		e := ErrNoRows{}
		if errors.As(err, &e) {
			fmt.Println(err)
			return
		}
		// 可能存在其他错误
	}
	fmt.Println(ret)
}

func GetUsers() ([]int, error) {
	ret, err := query()
	if err != nil {
		return nil, errors.Wrap(err, "未找到用户")
	}
	return ret, nil
}

func query() ([]int, error) {
	return nil, ErrNoRows{Msg: "no rows"}
}

type ErrNoRows struct {
	error
	Msg string
}

func (q ErrNoRows) Error() string {
	return q.Msg
}
