package main

var q = `
如果一个 sql 查询返回了一个不存在的错误，这个错误要怎么处理
`

type QueryError struct {
	Msg string
}

func (q QueryError) Error() string {
	return q.Msg
}

func query() error {
	return QueryError{Msg: "查询不存在"}
}

func handle() {
	err := query()
	if err != nil {

	}
}
