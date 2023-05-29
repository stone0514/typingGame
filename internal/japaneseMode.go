package internal

import (
	"bufio"
	"fmt"
	"github.com/stone0514/typingGame/kvs"
	"github.com/stone0514/typingGame/pkg"
	"os"
	"strings"
	"time"
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
 *               continue processing when retrying [any keys]
 *               display top10 scores in a Ranking [r]
 *               if not the process ends           [e]
 *
 *-------------------------------------------------------------
 */
func JapaneseMode() {

	inc := 0

	r := kvs.NewRedisClient()

	//parallel processineg of the read and input string in the channel
	ch := pkg.Input(os.Stdin)

	//create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- Enter your username ( up to 16 characters ) ---")
	//read the user input for the username
	scanner.Scan()
	username := scanner.Text()

	//truncate the username if it is longer than 16 characters
	if len(username) > 16 {
		username = username[:16]
	}

	//if the username is empty or only contains whitespace, set adefault
	if len(strings.TrimSpace(username)) == 0 {
		username = fmt.Sprintf("user-Ja-%02d", inc)
	}

	for i := 3; i >= 1; i-- {
		fmt.Println("--- ", i, " ---")
		time.Sleep(time.Second)
	}

	fmt.Println("\n--- Start ---")

	for {

		inc++

		//create timeLimit
		timeLimit := time.After(time.Duration(JaTime) * time.Second)

		fmt.Println("\n--- Go!JapaneseMode!(timelimit:", JaTime, "sec) ---")
		fmt.Println("--- 表示される単語を入力してください ---")
		fmt.Println("--- 制限時間内になるべく多くの単語を入力してください ---")

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
					fmt.Printf("\x1b[32m%s\x1b[0m\n", "----- Correct! -----")
				} else {
					fmt.Printf("\x1b[31m%s\x1b[0m\n", "----- Failure! -----")
				}
				//processing ends when the time limit
			case <-timeLimit:
				fmt.Printf("\n----- time up -----\n")
				i = false
			}
		}

		fmt.Println("\n----- result -----")
		fmt.Printf("----- score: %d -----\n", score)
		//result
		jResult(score)

		//save the Score
		kvs.SaveJaScore(r, username, score)
		inc++

		fmt.Println("\n--- retry[any keys] or show Ranking[r] or exit[e] ---")

		retryFlg := <-ch

		//if the last character is "r" show Ranking
		if strings.HasSuffix(retryFlg, "r") {
			kvs.JaRanking(r)
			return
		}

		//if the last character is "e" process ends
		if strings.HasSuffix(retryFlg, "e") {
			fmt.Println("----- End -----")
			return
		}

		//in other cases execute process again
		fmt.Println("-- Ready --")

		// delay
		<-time.Tick(time.Second)

		continue
	}
}
