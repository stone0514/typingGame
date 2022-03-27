package internal

// CreateQuestion Document
/*
 *------------------------------------------------
 * @brief CreateQuestion
 *
 * @param
 *
 * @retrun string words1
 *         string words2(hiragana)
 *
 * @rules
 *
 *------------------------------------------------
 */
func CreateQuestion() (string, string) {

	q := Question()

	var k, v string

	for k, v := range q {
		return k, v
	}
	return k, v
}
