package internal

import (
	"strings"
)

func ArrayContains(arr []string, str string) (bool, string) {
	for _, v := range arr {
		if strings.Contains(str, v) {
			if v == "n" {
				v = v + "n"
			}
			return true, v
		}
	}
	return false, ""
}
