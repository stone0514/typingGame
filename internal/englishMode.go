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

// EnglishMode Document
/*
 *-------------------------------------------------------------------
 * @brief EnglishMode
 *
 * @param[in]       string   internal.ChoiceGameMode *
 * @return[want]
 *
 * @rules           after calling the internal.ChoiceGameMode
 *                  selected 2 in ChoiceGameMode
 *
 *                  continue processing when retrying [any keys]
 *                  display top10 scores in a Ranking [r]
 *                  if not the process ends           [e]
 *
 *-------------------------------------------------------------------
 */
func EnglishMode() {

	inc := 0

	r := kvs.NewRedisClient()

	//parallel processing of the read and input string in the channel
	ch := pkg.Input(os.Stdin)

	//create a scanner to read user unput
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- Enter your username ( up to 16 characters ) ---")

	//read the user input for the username
	scanner.Scan()
	username := scanner.Text()

	//truncate the username if it is longer than 16 characters
	if len(username) > 16 {
		username = username[:16]
	}

	//if the username is empty or only contains whitespace, set a default
	if len(strings.TrimSpace(username)) == 0 {
		username = fmt.Sprintf("user-En-%02d", inc)
		inc++
	}

	for i := 3; i >= 1; i-- {
		fmt.Println("--- ", i, " ---")
		time.Sleep(time.Second)
	}

	fmt.Println("\n--- Start ---")

	for {

		inc++

		//create timeLimit
		timeLimit := time.After(time.Duration(EnTime) * time.Second)

		words := pkg.ReadLine(En)
		pkg.Shuffle(words)

		fmt.Println("\n--- Go!EnglishMode!(timelimit:", EnTime, "sec) ---")
		fmt.Println("--- 表示される単語を入力してください ---")
		fmt.Println("--- 制限時間内になるべく多くの単語を入力してください ---")

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
					fmt.Printf("\x1b[32m%s\x1b[0m\n", "----- Correct! -----")
				} else {
					fmt.Printf("\x1b[31m%s\x1b[0m\n", "----- Failure! -----")
				}
				//processing ends when the time limit
			case <-timeLimit:
				fmt.Println("\n----- time up -----")
				i = false
			}
		}

		fmt.Println("----- result -----")
		fmt.Printf("----- score: %d -----\n", score)
		//result
		Result(score)

		//save the Score
		kvs.SaveEnScore(r, username, score)
		inc++

		fmt.Println("\n--- retry[any keys] or show Ranking[r] or exit[e] ---")

		retryFlg := strings.TrimSpace(<-ch)

		//if the last character is "r" show Ranking
		if strings.HasSuffix(retryFlg, "r") {
			kvs.EnRanking(r)
			return
		}

		//if the last character is "e" process ends
		if strings.HasSuffix(retryFlg, "e") {
			fmt.Print("----- End -----")
			return
		}

		//in other cases execute process again
		fmt.Print("-- Ready --\n")

		// delay
		<-time.Tick(time.Second)
		continue
	}
}
