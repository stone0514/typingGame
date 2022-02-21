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

	// output question & ngram(uni-gram)
	for k, v := range m {
		// japanese
		fmt.Println("k :", k)
		// hiragana forAnalysis
		fmt.Println("v: ", v)
		tgtText := v
		fmt.Println("tgtText: ", tgtText)
		uni, err := ngram(tgtText, 1)
		if err != nil {
			fmt.Println("error : ", err)
		}
		// parse result
		fmt.Println("uni: ", uni)
	}

	// create dictionary
	dictionary := map[string][]string{
		"あ": {"a"},
		"い": {"i", "yi"},
	}
	fmt.Println("dictionary: ", dictionary)
	fmt.Println("dictionary test: ", dictionary["あ"])
	fmt.Println("dictionary test2: ", dictionary["い"])

	// dictionary check
	for k, v := range dictionary {
		fmt.Println("k: ", k)
		fmt.Println("v: ", v)
	}

	// ngram sample(uni-gram)
	// tgtText := "本当に今日は良い天気ですね"
	tgtText := "あいあいあい"
	uni, err := ngram(tgtText, 1)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("uni: ", uni)

	// parse uni-gram = dictionary(v hiragana)
	ret := make([]string, 0, 0)
	for i := 0; len(uni) > i; i++ {
		for k, v := range dictionary {
			if k == uni[i] {
				fmt.Println("i route: ", v)
				fmt.Println("k: ", k)
				fmt.Println("v: ", v)
				ret = append(ret, v...)
			}
		}
	}
	fmt.Println("ret: ", ret)
	fmt.Println("out")
}

// create ngram for testFunc
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
