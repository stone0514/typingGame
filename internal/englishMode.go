package internal

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/stone0514/typingGame/pkg"
)

// EnglishMode Document
/*
 *-------------------------------------------------------------------
 * @brief EnglishMode
 *
 * @param[in]       string   internal.ChoiceGameMode
 *
 * @return[want]
 *
 * @rules           after calling the internal.ChoiceGameMode
 *                  selected 2 in ChoiceGameMode
 *
 *                  continue processing when retrying
 *                  if not the process ends
 *
 *-------------------------------------------------------------------
 */
func EnglishMode(gameMode string) {

	//parallel processing of the read and input string in the channel
	ch := pkg.Input(os.Stdin)
retry:
	//create timeLimit
	t := EnTime

	timeLimit := time.After(time.Duration(t) * time.Second)

	words := pkg.ReadLine(gameMode)
	pkg.Shuffle(words)

	fmt.Println("\n---Go!EnglishMode!(timelimit:", t, "sec)---")

	score := 0
	num := 1

	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Printf("[%d] %s\n", num, question)
		select {
		case x := <-ch:
			//correct/wrong judgement
			if question == x {
				score++
				num++
				fmt.Printf("\x1b[32m%s\x1b[0m\n", "-----Correct!-----")
			} else {
				fmt.Printf("\x1b[31m%s\x1b[0m\n", "-----Failure!-----")
			}
			//processing ends when the time limit
		case <-timeLimit:
			fmt.Printf("\n-----time up -----\n")
			i = false
		}
	}

	fmt.Printf("\n----- result -----")
	fmt.Printf("\n----- score: %d -----\n\n", score)
	//result
	Result(score)

	fmt.Print("\n--- retry[any keys] or exit[e] ---\n")

	retryFlg := <-ch

	//if the last character is "e" process ends
	if strings.HasSuffix(retryFlg, "e") {
		fmt.Print("----- End -----")
		//in other cases execute process again
	} else {
		fmt.Print("-- Ready --")
		time.Sleep(time.Second * 1)
		goto retry
	}
}
