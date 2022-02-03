package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLine Document
/*
 *------------------------------------------------------
 * @brief    ReadLine
 *           Read one line and append Slice
 *
 * @param[1] string need filePath
 *
 * @return   []string words in Slice
 *
 * @rules    filePath = input Path from the execution Directory onward
 *------------------------------------------------------
 */
func ReadLine(filePath string) []string {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "filePath %s could not read", filePath, err)
		os.Exit(1)
	}

	defer f.Close()

	line := make([]string, 0, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "file %s scan error %v\n", filePath, err)
	}

	return line
}
