package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stone0514/typingGame/internal"
	"github.com/stone0514/typingGame/pkg"
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

	//create timeLimit
	t := int(75)
	pkg.Init(t)

	timeLimit := time.After(time.Duration(t) * time.Second)

	//mode select
	fmt.Printf("\nchoice play mode\n")
	fmt.Printf("1:English, 2:Japanese\n")
	gameMode := internal.ChoiceGameMode(os.Stdin)

	//read word & shuffle
	words := pkg.ReadLine(gameMode)
	pkg.Shuffle(words)

	fmt.Println("\nStart!(timelimit:", t, "sec)")

	score := 0
	num := 1
	//parallel processing of the read and input string in the ch
	ch := pkg.Input(os.Stdin)
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Printf("[%d] %s\n", num, question)

		select {
		case x := <-ch:
			// correct / wrong judgment
			if question == x {
				score++
				num++
			} else {
				fmt.Print("-----miss!retype!-----\n")
			}
			// processing ends when the time limit
		case <-timeLimit:
			fmt.Printf("\n----- time up -----\n")
			i = false
		}
	}

	fmt.Printf("\n----- result -----")
	fmt.Printf("\n----- score: %d -----\n\n", score)
	// result
	internal.Result(score)

}
