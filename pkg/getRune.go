package pkg

// GetRune Document
/*
 *---------------------------------------------------
 * @brief        GetRune
 *
 * @param[in]    string
 * @param[in]    int
 *
 * @return[want] rune
 *
 * @rules
 *
 *
 *---------------------------------------------------
 */
func GetRune(s string, i int) rune {
	//make s array
	rs := []rune(s)
	return rs[i]
}
