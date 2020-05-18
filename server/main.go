package main

import (
	"log"
	"net/http"

	"server/views"
)

func main() {
	// 静态文件夹
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// 路由
	http.HandleFunc("/handle", views.Handle)
	http.HandleFunc("/heihei", views.Heihei)
	// http.HandleFunc("/", views.Settings)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
