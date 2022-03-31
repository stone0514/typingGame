package internal

import "fmt"

// Result Document
/*
 *---------------------------------------------------
 * @brief     Result
 *            showResult
 *
 * @param[in] int score
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
		fmt.Println("----- you noob:) -----")
	case n >= 16 && n <= 25:
		fmt.Println("----- your beginner -----")
	case n >= 25 && n <= 37:
		fmt.Println("------ your class is expart -----")
	case n >= 38:
		fmt.Println("----- Awesome! -----")

	default:
		fmt.Println("don't give up")
	}
}
