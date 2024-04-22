package main

import (
	"fmt"
	"os"

	"github.com/mari-dotworld/passgen/hashpass"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("see help")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "hash":
		pass := os.Args[2]
		if pass == "" {
			fmt.Println("see help")
			return
		}
		hashPass := hashpass.HashPassword(pass)
		fmt.Println("hash : ", hashPass)
	case "compare":
		hash := os.Args[2]
		pass := os.Args[3]
		if hash == "" || pass == "" {
			fmt.Println("see help")
			return
		}
		verify := hashpass.ComparePassword(pass, hash)
		if verify {
			fmt.Println("password matched")
		} else {
			fmt.Println("password not matched")
		}
	case "help":
		fmt.Println("list of commands : ")
		fmt.Println("./passgen hash [password]")
		fmt.Println("./passgen compare [hash] [password]")
	}
}
