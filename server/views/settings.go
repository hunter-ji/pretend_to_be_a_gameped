package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	// 处理设置页面表单
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/settings.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "You are loged in !!!")
	}
}
