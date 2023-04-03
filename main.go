package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	var text string
	var arg string

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()

		if arg == "--print" {
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		os.Exit(1)
	}

	err := clipboard.WriteAll(text)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ command output copied to clipboad")
}
