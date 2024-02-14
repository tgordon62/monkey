package main

import (
	"fmt"
	"monkey/filearg"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		readEvalPrintLoop(user)
	} else {
		filearg.Start(os.Args[1])
	}

}

func readEvalPrintLoop(user *user.User) {
	fmt.Printf("Hello %s! This is the Monkey language console!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
