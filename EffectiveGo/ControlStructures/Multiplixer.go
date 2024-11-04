package main

import (
	"fmt"
	"time"
)

func Multiplixer() {
	// iki kanal tanımladık.
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "kanal 1 verisi(Kameralardaki insan yüzleri)"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "parmak izi verisi"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1) //kanal birden veri alırsa burası çalışır.
		case msg2 := <-channel2:
			fmt.Println(msg2) //kanal ikiden veri gelirse bura çalışır.
		}
	}
}
