package main

import "fmt"

//Go dilindeki for döngüsü, C dilindekine benzer.
//for ve while döngülerini birleştirir. do-while döngüsü goda yok.
//forda 3 farklı biçimi vardır. Bunlardan yalnızca birinde noktalı virgül kıllanılır.
//tekrar edilen işlemlerde, sayma işlemlerinde, dizi, liste ve dizi üzerine işlemlerde,zamanlayıcı işlemlerinde.
//Koşul sonuçlarını belirtmede, tekrar edilen görevlerin yönetiminde işimize yarayacak.

//Like a C for
//for init; condition; post{}

func forStruct() {
	//Like a C for
	//for init; condition; post{}
	// for deger := 0; deger < 10; deger += 2 {
	// 	println(deger)
	// } //baslangıç değerimiz 0 dır 10 dan küçük oluncaya kadar 2 şer 2 şer artıtrır. 8 de durur.

	//Like a C While

	// deger := 0
	// for deger < 10 {
	// 	fmt.Println(deger)
	// 	deger += 2
	// }

	//Like a C for(;;)
	for {
		var input string
		fmt.Println("Birşeyler yazın çıkmak için exit yazın")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println("Çıkılıyor...")
			break // Döngüden çık
		}

		fmt.Println("Girdiğiniz:", input)
	}
}
