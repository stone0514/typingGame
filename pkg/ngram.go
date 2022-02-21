package pkg

import (
	"errors"
	"strings"
)

// Ngram Document
/*
 *----------------------------------------------------------------------
 * @brief    Ngram
 *
 * @param    P1: string     targetText
 *           P2: int        n(word count)
 *
 * @return   P1: []string
 *           P2: error
 *
 * @rules    +create ngram
 *           +param:P2 int(1) = uni-gram
 *           +param:P2 int(2) = bi-gram
 *           +param:P2 int(3) = tri-gram
 *
 *----------------------------------------------------------------------
 */

func Ngram(tgtText string, n int) ([]string, error) {
	var ngrams []string
	spltText := strings.Split(tgtText, "")
	if len(spltText) < n {
		err := errors.New("Error: Input string's length is less than n value")
		return nil, err
	}
	for i := 0; i < len(spltText)-n+1; i++ {
		ngrams = append(ngrams, strings.Join(spltText[i:i+n], ""))
	}
	return ngrams, nil
}
