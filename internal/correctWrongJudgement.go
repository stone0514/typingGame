package internal

import (
	"strings"
)

// CorrectWrongJudgement Document
/*
 *-----------------------------------------------------------
 * @brief CorrectWrongJudgement
 *
 * @param[in]    map[string][]string   m
 * @param[in]    []string              arr
 * @param[in]    string                str
 *
 * @return[want] bool
 * @return[want] string
 *
 * @rules        m is create by GenerationInputPattern
 *               arr is create by SeparateWords
 *
 *-----------------------------------------------------------
 */
func CorrectWrongJudgement(m map[string][]string, arr []string, str string) (bool, string) {

	//create judgeList
	judgeList := map[int][]string{}

	//loop for length of the array
	for i := 0; i < len(arr); i++ {
		for k, v := range m {
			if k == arr[i] {
				judgeList[i] = append(judgeList[i], v...)
				b, mStr := ArrayContains(judgeList[i], str)
				if b {
					//delete matching letter
					str = strings.Replace(str, mStr, "", 1)
				} else {
					return false, str
				}
			}
		}
	}
	return true, str
}
