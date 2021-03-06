package internal

import "fmt"

// Result Document
/*
 *---------------------------------------------------
 * @brief        Result
 *
 * @param[in]    int score
 *
 * @return[want]
 *
 * @rules
 *
 *---------------------------------------------------
 */
func Result(score int) {

	n := score
	switch {
	case n <= 15:
		fmt.Println("----- don't give up  -----")
	case n >= 16 && n <= 28:
		fmt.Println("----- your beginner -----")
	case n >= 29 && n <= 39:
		fmt.Println("------ your expart! -----")
	case n >= 40:
		fmt.Println("----- Awesome! -----")

	default:
		fmt.Println("don't give up")
	}
}
