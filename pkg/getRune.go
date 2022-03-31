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
 *  convert from string to rune
 *
 *---------------------------------------------------
 */
func GetRune(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}
