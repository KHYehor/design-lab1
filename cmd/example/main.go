package main

import (
	"errors"
	"fmt"
	"os"

	lab1 "github.com/KHYehor/lab1"
)

func exitWithErr(message string) {
	fmt.Println(errors.New(message))
	os.Exit(1)
}

func main() {
	if len(os.Args[0:]) < 2 {
		exitWithErr("must be addition arg")
	} else if len(os.Args[0:]) > 2 {
		exitWithErr(`you command must be like: go run /path/to/program/main.go "your postfix string"`)
	}

	postfixString := os.Args[1]

	prefixString, err := lab1.PostfixToPrefix(postfixString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Your prefix expression:", prefixString)
	return
}
