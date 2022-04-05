package internal

import "fmt"

// jResult Document
/*
 *---------------------------------------------------
 * @brief        jResult
 *
 * @param[in]    int score
 *
 * @return[want]
 *
 * @rules
 *
 *---------------------------------------------------
 */
func jResult(score int) {

	n := score
	switch {
	case n <= 149:
		fmt.Println("----- don't give up  -----")
	case n >= 150 && n <= 175:
		fmt.Println("----- your beginner -----")
	case n >= 176 && n <= 199:
		fmt.Println("------ your expart! -----")
	case n >= 200:
		fmt.Println("----- Awesome! -----")

	default:
		fmt.Println("don't give up")
	}
}
