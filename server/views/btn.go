package views

import (
	"encoding/json"
	"net/http"
)

type btnMap struct {
	BtnLeft1  string `json:"btnleft1"`
	BtnLeft2  string `json:"btnleft2"`
	BtnLeft3  string `json:"btnleft3"`
	BtnLeft4  string `json:"btnleft4"`
	BtnRight1 string `json:"btnright1"`
	BtnRight2 string `json:"btnright2"`
	BtnRight3 string `json:"btnright3"`
	BtnRight4 string `json:"btnright4"`
}

type Res struct {
	Code int `json:"code"`
}

func Btn(w http.ResponseWriter, r *http.Request) {
	var b btnMap
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	

	w.Header().Set("Content-Type", "application/json")
	res := Res{
		Code: 20000,
	}
	json.NewEncoder(w).Encode(res)
}
