package pkg

import (
	"bufio"
	"io"
	"strings"
)

// Input Document
/*
 *-------------------------------------------------
 * @brief        Input
 *               InputReader
 *
 * @param[in]    io.Reader
 *
 * @return[want] chan string
 *
 * @rules
 *
 *-------------------------------------------------
 */
func Input(r io.Reader) <-chan string {
	// create ch
	ch := make(chan string)
	// go routine
	go func() {
		// read the typed line
		reader := bufio.NewReader(r)
		for {
			word, err := reader.ReadString('\n')
			if err != nil {
				close(ch)
				return
			}
			word = strings.TrimSpace(word) //空白文字を削除
			ch <- word
		}
	}()
	return ch
}
