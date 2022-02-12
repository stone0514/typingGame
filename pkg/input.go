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
	// create chan
	channel := make(chan string)
	// go routine
	go func() {
		// read the typed line
		strings := bufio.NewScanner(stdin)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}
