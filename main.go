package main

import (
	"Nova/repel"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s This is Nova Programming Language!\n",
		user.Username)
	fmt.Printf("Feel Free to roam in the Nebula\nEnter quit to exit\n")
	repel.Start(os.Stdin, os.Stdout)
}
