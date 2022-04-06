package internal

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/stone0514/typingGame/pkg"
)

// JapaneseMode Document
/*
 *-------------------------------------------------------------
 * @brief JapaneseMode
 *
 * @param[in]    string   internal.ChoiceGameMode
 *
 * @return[want]
 *
 * @rules        after calling the internal.ChoiceGameMode
 *               selected 2 in ChoiceGameMode
 *
 *               continue processing when retrying
 *               if not the process ends
 *
 *-------------------------------------------------------------
 */
func JapaneseMode() {

	//parallel processing of the read and input string in the channel
	ch := pkg.Input(os.Stdin)
retry:
	//create timeLimit
	t := JaTime
	timeLimit := time.After(time.Duration(t) * time.Second)

	fmt.Println("\n---Go!JapaneseMode!(timelimit:", t, "sec)---")

	score := 0
	num := 1

	for i := true; i; {

		k, v := Question()
		fmt.Printf("[%d] %s\n", num, k)
		fmt.Printf("[+] %s\n", v)

		ret := SeparateWords(v)

		concatList := GenerationInputPattern(ret)
		select {
		case x := <-ch:
			//correct/wrong judgement
			j, s := CorrectWrongJudgement(concatList, ret, x)
			if j && s == "" {
				score = score + len(ret)
				num++
				fmt.Print("-----Correct!-----\n")
			} else {
				fmt.Print("-----Failure!-----\n")
			}
			//processing ends when the time limit
		case <-timeLimit:
			fmt.Printf("\n----- time up -----\n")
			i = false
		}
	}

	fmt.Printf("\n----- result -----")
	fmt.Printf("\n----- score: %d -----\n", score)
	jResult(score)

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
