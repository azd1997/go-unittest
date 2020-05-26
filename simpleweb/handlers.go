package main

import (
	"fmt"
	"net/http"
)

// 查看打招呼的hello，同样打招呼回去，做一个回显
func handleHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

// 发邮件。请求中包含邮件接收人信息，服务器替客户端发送这封邮件
func handleMail(w http.ResponseWriter, r *http.Request) {

}

// 查数据库。
func handleQuery(w http.ResponseWriter, r *http.Request) {

}


var f = ttt.Func

type ttt struct {
}

func (t *ttt) Func() {
}

func F() {
	f = func(t ttt) {

	}
}