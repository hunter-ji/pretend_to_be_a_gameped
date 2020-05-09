package views

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	absPath, _ := filepath.Abs("templates/handle.html")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("rad file err:", err.Error())
		return
	}
	fmt.Fprintf(w, string(data))
}
