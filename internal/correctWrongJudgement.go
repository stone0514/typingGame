package internal

import (
	"strings"
)

// CorrectWrongJudgement Document
/*
 *-----------------------------------------------------------
 * @brief CorrectWrongJudgement
 *
 * @param[in]    map[string][]string
 * @param[in]    []string
 * @param[in]    string
 *
 * @return[want] bool
 * @return[want] string
 *
 * @rules
 *
 *-----------------------------------------------------------
 */
func CorrectWrongJudgement(m map[string][]string, arr []string, str string) (bool, string) {

	judgeList := map[int][]string{}

	for i := 0; i < len(arr); i++ {
		for k, v := range m {
			if k == arr[i] {
				judgeList[i] = append(judgeList[i], v...)
				b, mStr := ArrayContains(judgeList[i], str)
				if b {
					str = strings.Replace(str, mStr, "", 1)
				} else {
					return false, str
				}
			}
		}
	}
	return true, str
}
