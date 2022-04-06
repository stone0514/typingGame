package internal

import (
	"unicode/utf8"

	"github.com/stone0514/typingGame/pkg"
)

// SeparateWords Document
/*
 *-----------------------------------------------------
 * @brief
 *
 * @param[in]    string    words
 *
 * @return[want] []string
 *
 * @rules        words must hiragana
 *
 *-----------------------------------------------------
 */
func SeparateWords(words string) []string {

	var (
		uni, bi string
		ret     []string
	)

	i := 0
	d := pkg.Dictionary()

	//word length loop
	for i < utf8.RuneCountInString(words) {
		//stores i character & cast rune to string
		uni = string(pkg.GetRune(words, i))
		if i+1 < utf8.RuneCountInString(words) {
			//stores i & i+1 character
			bi = string(pkg.GetRune(words, i)) + string(pkg.GetRune(words, i+1))
		} else {
			bi = ""
		}
		if _, ok := d[bi]; ok {
			i += 2
			ret = append(ret, bi)
		} else {
			i++
			ret = append(ret, uni)
		}
	}
	return ret
}
