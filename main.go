package main

import (
	"fmt"
	"os"
	"time"
	"typingGame/internal"
	"typingGame/pkg"
)

// main
/*
 *--------------------------------------------------------------
 * @brief main
 *        typing game
 *
 * @param
 *
 * @return
 *
 * @rules
 *
 *--------------------------------------------------------------
 */
func main() {

	t := int(75)
	pkg.Init(t)

	timeLimit := time.After(time.Duration(t) * time.Second)

	fmt.Printf("\nchoice play mode\n")
	fmt.Printf("1:English, 2:Japanese\n")
	gameMode := internal.ChoiceGameMode(os.Stdin)

	words := pkg.ReadLine(gameMode)
	pkg.Shuffle(words)

	fmt.Println("\nStart!(timelimit:", t, "sec)")

	score := 0
	num := 1
	ch := pkg.Input(os.Stdin)
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Printf("[%d] %s\n", num, question)

		select {
		case x := <-ch:
			if question == x {
				score++
				num++
			} else {
				fmt.Print("-----miss!retype!-----\n")
			}

		case <-timeLimit:
			fmt.Printf("\n----- time up -----\n")
			i = false
		}
	}

	fmt.Printf("\n----- result -----")
	fmt.Printf("\n----- score: %d -----\n\n", score)

	internal.Result(score)

}
