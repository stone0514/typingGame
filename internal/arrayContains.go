package internal

import (
	"strings"
)

// ArrayContains Document
/*
 *--------------------------------------------------------------
 * @brief ArrayContains
 *
 * @param[in]    []string
 * @param[in]    string
 *
 * @return[want] bool
 * @return[want] string
 *
 * @rules
 *
 *--------------------------------------------------------------
 */
func ArrayContains(arr []string, str string) (bool, string) {
	for _, v := range arr {
		if strings.Contains(str, v) {
			if v == "n" {
				if strings.Contains(str, "nn") {
					v = "nn"
				}
				if strings.Contains(str, "n'") {
					v = "n'"
				}
				if strings.Contains(str, "xn") {
					v = "xn"
				}
			}
			return true, v
		}
	}
	return false, ""
}
