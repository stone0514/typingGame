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

	ch := pkg.Input(os.Stdin)
retry:
	t := int(90)
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
			j, s := CorrectWrongJudgement(concatList, ret, x)
			if j && s == "" {
				score = score + len(ret)
				num++
				fmt.Print("-----Correct!-----\n")
			} else {
				fmt.Print("-----Failure!-----\n")
			}
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

	if strings.HasSuffix(retryFlg, "e") {
		fmt.Print("----- End -----")
	} else {
		fmt.Print("-- Ready --")
		time.Sleep(time.Second * 1)
		goto retry
	}
}
