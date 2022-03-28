package internal

import (
	"strings"
)

func ArrayContains(arr []string, str string) (bool, string) {
	for _, v := range arr {
		if strings.Contains(str, v) {
			return true, v
		}
	}
	return false, ""
}
