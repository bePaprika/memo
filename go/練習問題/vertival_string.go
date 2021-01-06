package main

import (
	"fmt"
)

func main() {
	// 変数定義
	var message string
	var c byte

	// 入力受け取り 
	fmt.Printf("Enter a string :")
	fmt.Scan(&message)
	fmt.Println(":")

	// 表示
	for i := 0; i<len(message); i++{
		c = message[i]
		fmt.Printf("%c\n",c)
	}

	// 表示
	for _, c := range message {
		fmt.Printf("%c\n",c)
	}
}