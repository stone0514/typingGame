package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/stone0514/typingGame/internal"
	"github.com/stone0514/typingGame/pkg"
)

func main() {

	/*
		t := int(75)
		pkg.Init(t)

		timeLimit := time.After(time.Duration(t) * time.Second)
	*/

	q := internal.Question()
	//d := pkg.Dictionary()

	//score := 0
	num := 1

	//ch := pkg.Input(os.Stdin)

	for true {
		question := q //[score]
		for k, v := range question {
			fmt.Printf("[%d] %s\n", num, k)
			fmt.Printf("[%d] %s\n", num, v)

			i := 0
			var (
				uni string
				bi  string
			)

			for i < utf8.RuneCountInString(v) {
				uni = string(pkg.GetRune(v, i))
				fmt.Println("uni : ", uni)
				if i+1 < utf8.RuneCountInString(v) {
					bi = string(pkg.GetRune(v, i))
					fmt.Println("bi : ", bi)
				} else {
					bi = ""
				}
				return
			}
		}
	}
}
