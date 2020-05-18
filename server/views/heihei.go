package views

import (
	"encoding/json"
	"log"
	"net/http"
)

type btnMap struct {
	code string
	data string
}

func Heihei(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t btnMap
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.data)

}
