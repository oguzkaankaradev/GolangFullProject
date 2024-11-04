package main

import "fmt"

func main() {

	var nargileKafeYas = 18
	var CustomerAge int
	fmt.Println("Lütfen yaşınızı giriniz")
	fmt.Scan(&CustomerAge)

	if CustomerAge >= nargileKafeYas {
		fmt.Println("Hoşgeldiniz")
	} else {
		fmt.Println("velinizle geliniz.")
	}
}
