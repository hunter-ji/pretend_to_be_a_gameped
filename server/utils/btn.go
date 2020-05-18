package utils

import (
	"github.com/go-vgo/robotgo"
)

func oneBtn(key string) {
	robotgo.KeyTap(key)
}

func twoBtn(key, key2 string) {
	robotgo.KeyTap(key, key2)
}

func toggleBtn(key string, module bool) {
	if module {
		robotgo.KeyToggle(key, "down")
	} else {
		robotgo.KeyToggle(key, "up")
	}

}
