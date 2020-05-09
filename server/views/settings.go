package views

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	// 处理设置页面表单
	r.ParseForm()
	if r.Method == "GET" {
		absPath, _ := filepath.Abs("templates/settings.html")
		data, err := ioutil.ReadFile(absPath)
		if err != nil {
			fmt.Println("rad file err:", err.Error())
			return
		}
		fmt.Fprintf(w, string(data))
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprintf(w, "You are loged in !!!")
	}
}
