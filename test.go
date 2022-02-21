package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	// questionList
	m := map[string]string{"朝": "あさ", "昼": "ひる", "夜": "よる"}

	fmt.Println(m)

	// create question & parse
	for k, v := range m {
		// japanese
		fmt.Println(k)
		// hiragana forAnalysis
		fmt.Println(v)
		tgtText := v
		fmt.Println(tgtText)
		bigrams, err := ngram(tgtText, 1)
		if err != nil {
			fmt.Println(err)
		}
		// parse result
		fmt.Println("bigrams: ", bigrams)
	}

	// create dictionary
	dictionary := map[string][]string{
		"あ": {"a"},
		"い": {"i", "yi"},
	}
	fmt.Println(dictionary)
	fmt.Println(dictionary["あ"])
	fmt.Println(dictionary["い"])

	for k, v := range dictionary {
		fmt.Println(k)
		fmt.Println(v)
	}

	// parse sample
	//tgtText := "本当に今日は良い天気ですね"
	tgtText := "あいあいあい"
	bigrams, err := ngram(tgtText, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bigrams: ", bigrams)

	// correct/wrong judgment
	for i := 0; len(bigrams) > i; i++ {
		for k, v := range dictionary {
			if k == bigrams[i] {
				fmt.Println("i route: ", v)
				fmt.Println("k: ", k)
				fmt.Println("v: ", v)
			}
		}

	}
	fmt.Println("out")

}

// parse hiragana func
func ngram(tgtText string, n int) ([]string, error) {
	spltText := strings.Split(tgtText, "")
	var ngrams []string
	if len(spltText) < n {
		err := errors.New("Error: Input string's length is less than n value")
		return nil, err
	}
	for i := 0; i < len(spltText)-n+1; i++ {
		ngrams = append(ngrams, strings.Join(spltText[i:i+n], ""))
	}
	return ngrams, nil
}
