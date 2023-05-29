package main

import (
	"fmt"
	"github.com/stone0514/typingGame/internal"
	"github.com/stone0514/typingGame/kvs"
	"log"
	"os"
	"os/exec"
)

// main Document
/*
 *--------------------------------------------------------------
 * @brief main
 *        typing game
 *
 * @param
 *
 * @return
 *
 * @rules this is typingGame
 *        selectable from English or JapaneseMode
 *        continue processing when retrying [any keys]
 *        display top10 scores in a Ranking [r]
 *        if not the process ends           [e]
 *
 *--------------------------------------------------------------
 */
func main() {

	//start redis server
	cmd := exec.Command("redis-server")
	//start the redis server in the background
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis server is running in the background")

	//wait for the program to exit, then stop the redis server
	defer func() {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Printf("Failed to kill redis server %v\n", err)
		}
	}()

	//connection redis server
	kvs.Client()

	//mode select
	fmt.Println("----- choice play mode -----")
	fmt.Println("--- 1:English, 2:Japanese ---")

	//wait for input
	gameMode, err := internal.ChoiceGameMode(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch gameMode {
	//input = "1"
	case internal.En:
		internal.EnglishMode()
	//input = "2"
	case internal.Ja:
		internal.JapaneseMode()
	//other inputs
	default:
		internal.JapaneseMode()
	}
}
