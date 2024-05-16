package main

import "fmt"

func main() {
	//Variable birinci yöntem
	// var greetingText string
	// greetingText = "Hello World!!!!!"

	//Variable ikinci yöntem
	// var greetingText string = "Hello World!!!!!"

	greetingText := "Hello World!!!!!"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5 //17 + 5 oldu.

	fmt.Println(greetingText)
	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = luckyNumber * 3
	fmt.Println(evenMoreLuckyNumber)

}
