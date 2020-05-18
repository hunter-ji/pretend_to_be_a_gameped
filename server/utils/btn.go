package utils

import (
	"github.com/go-vgo/robotgo"
)

func OneBtn(key string) {
	robotgo.KeyTap(key)
}

func TwoBtn(key, key2 string) {
	robotgo.KeyTap(key, key2)
}

func ToggleBtn(key string, module bool) {
	if module {
		robotgo.KeyToggle(key, "down")
	} else {
		robotgo.KeyToggle(key, "up")
	}

}
