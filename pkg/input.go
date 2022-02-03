package pkg

import (
	"bufio"
	"io"
)

// Input Document
/*
 *-------------------------------------------------
 * @brief        Input
 *               InputReader
 *
 * @param[in] P1 io.Reader
 *
 * @return
 *
 * @rules
 *
 *-------------------------------------------------
 */
func Input(stdin io.Reader) <-chan string {
	channel := make(chan string)
	go func() {
		strings := bufio.NewScanner(stdin)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}
