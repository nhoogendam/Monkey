package main

import (
	"Monkey/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is my monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}