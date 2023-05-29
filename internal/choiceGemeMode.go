package internal

import (
	"bufio"
	"errors"
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
func ChoiceGameMode(stdin io.Reader) (string, error) {
	val := bufio.NewScanner(stdin)
	var gameMode string

	if !val.Scan() {
		return "", errors.New("unable to scan input")
	}

	switch val.Text() {
	case "1":
		gameMode = En
	case "2":
		gameMode = Ja
	// other inputs
	default:
		gameMode = Ja
	}

	if err := val.Err(); err != nil {
		return "", err
	}

	return gameMode, nil
}
