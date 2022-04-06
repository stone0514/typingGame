package main

import (
	"fmt"
	"os"

	"github.com/stone0514/typingGame/internal"
)

// main Document
/*
 *--------------------------------------------------------------
 * @brief main
 *        typing game
 *
 * @param
 *
 * @return
 *
 * @rules this is typingGame
 *        selectable from English or JapaneseMode
 *        continue processing when retrying
 *        if not the process ends
 *
 *--------------------------------------------------------------
 */
func main() {

	//mode select
	fmt.Printf("\n-----choice play mode-----\n")
	fmt.Printf("-----select game start-----\n")
	fmt.Printf("---1:English, 2:Japanese---\n")
	//wait for input
	gameMode := internal.ChoiceGameMode(os.Stdin)

	switch gameMode {
	//input = "1"
	case internal.En:
		internal.EnglishMode(gameMode)
	//input = "2"
	case internal.Ja:
		internal.JapaneseMode()
	//other inputs
	default:
		internal.JapaneseMode()
	}
}
