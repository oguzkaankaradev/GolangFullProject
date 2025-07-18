package main

import "fmt"

//Go dilindeki Control Struct lar c diline benzemektedir.
//Fakat temel birkaç farklılık bulunmaktadır.
//Go da do veya while dögüsü yoktur
//Bunun yerine daha esnek for döngüsü kullanılır
//switch , c diline göre daha esnektir.
//if ve switch tıpkı for gibi, isteğe bağlı bir başlatma(initianization) ifadesi alabilir.
//break ve continue komutları label(Etiket) verilebilir.
//Bu sayede hangi döngünün kırılacağını veya devam edebiliecğine karar verebilirsin.
//Go da ayrıca
//Type switch(Tip konrolü yapan switch)
//select gibi yeni kontrol yapıları vardır.
//Söz diizmleri farklıdır.
//() parantez kullanılmaz, if, switch, for gibi yapılarda
//Kod blokları daima süslü parantez ile çevirilmek zorunda.

func main() {
	//User(ali,ahmet,mehmet)
	for User := 0; User < 50; User++ {
		fmt.Println(User) //İsim listesini tek tek dolaştı ekrana bastı.
	}

	//if yapısında parantez yok

	Age := 18
	if Age > 18 {
		fmt.Println("Yaşınız 18 den büyüktür.")
	}

	programStatus := "Start"
	switch programStatus {
	case "Start":
		fmt.Println("Start kodları çalışsın")
	case "Error":
		fmt.Println("Error kodlarında yapılacaklar çalışsın")
	default:
		fmt.Println("Default durumda yapılacaklar")
	_:
		fmt.Println("Diğer tüm durumlarda çalışacak kod")
	}

	//switch int

	switch num := number(); {

	case num > 0:
		fmt.Println("Pozitif")
	case num < 0:
		fmt.Println("negatif")
	default:
		fmt.Println("sıfıt")
	_:
		fmt.Println("Karmaşık sayılar.")

	}

	//select
	select {
	case msg := <-channel1:
		fmt.Println("Kanal 1 den geldi", msg)
	case msg := <-channel2:
		fmt.Println("Kanla 2 den geldi", msg)
	default:
		fmt.Println("Hiçbir kanaldan gelmedi.")
	}

	//break // Continue kullanımı
loop:
	for i := 0; i < 5; i++ {
		if i == 3 {
			break loop //loop diye bir etiket verdik etikete göre kırdık.

		}

		fmt.Println("i:", i)

	}

	//switch type
	var val interface{} = "merhaba"

	switch v := val.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("integer", v)
	default:
		fmt.Println("Bilinmeyen tür")
	}

	//Select üzerinden veri konrolü
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "Kanal birden gelen veri"

	select {
	case msg1 := <-ch1:
		fmt.Println("geldi::", msg1)
	case msg2 := <-ch2:
		fmt.Println("geldi:", msg2)
	default:
		fmt.Println("Hiçbir kanaldan veri gelmedi")
	}
}
