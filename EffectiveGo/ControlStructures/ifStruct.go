package main

import (
	"fmt"
	"os"
)

/*
İf  eğer koşul sağlanmış veya sağlanmamışsa duruma göre kontrol eder.
İf yapısının içerisinde return veya break gibi kontrol ifadeleri kullanılır.

return işlemin sonucunu içerisinde tutar, sen çağırana kadar bir sonuç yazmaz.

* if ve switch genellikle initialize(başlatma) ifadesi olrak kabul edildiği için
genellikle llokal bir değişkeni veya işlemi başalatmak için kullanılır.
*/
//nill null demektir.
func ExceptionHandlingExample() {
	//Chmod dosya izinlerini düzenlemek için kullanılır.
	//Bu kod ile sen dosya izinlerine müdahale edebilirsin.

	err := os.Chmod("Deneme.txt", 0644)
	if err != nil {
		fmt.Println("chmod işlemi başarısız:", err)
	} else {
		fmt.Println("dosya izinleri değiştirildi.")
	}
}

func OpenChmod() {
	f, err := os.OpenFile("Tanim.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("dosya açılmadı err", err)
		return
	}

	defer f.Close()

	err = f.Chmod(0666)
	if err != nil {
		fmt.Println("izin değiştirme hatası", err)
		return
	}

	_, err = f.WriteString("merhaba go")
	if err != nil {
		fmt.Println("yazma hatası", err)

		return
	}

	fmt.Println("işlemler başarıyla tamamlandı")
}
