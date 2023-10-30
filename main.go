package main

import  (
	"github.com/KoMaTop10/Monkey/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) == 2 {
		fileContent, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Err")
		}
		doScriptFile(string(fileContent))
		os.Exit(0)
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!! \n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin,os.Stdout)
}

func doScriptFile(fileContent string) {
	fmt.Printf(fileContent)
	fmt.Printf("\n")

	repl.domonkeyfile(fileContent, os.Stdout)
}