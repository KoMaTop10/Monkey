package main

import (
	"fmt"
	"os"
)

func main() {
	
	fileContent, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println(string(fileContent)) // ファイルの内容を出力
}