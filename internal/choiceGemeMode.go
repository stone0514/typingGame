package internal

import (
	"bufio"
	"io"
)

// ChoiceGameMode Document
/*
 *---------------------------------------------
 * @brief           ChoiceGameMode
 *                  Select GameMode
 *
 * @ param[in]      io.Reader
 *
 * @ return[want]   string filePath
 *
 * @rules           after receiving the input
 *                  input key 1: EnglishMode, 2: JapaneseMode
 *                  input other key 2: JapaneseMode
 *---------------------------------------------
 */
func ChoiceGameMode(stdin io.Reader) string {
	val := bufio.NewScanner(stdin)

	switch val.Scan() {
	case val.Text() == "1":
		return En
	case val.Text() == "2":
		return Ja
	// other inputs
	default:
		return Ja
	}
}
