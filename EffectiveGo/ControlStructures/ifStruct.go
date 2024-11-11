package main

import (
	"fmt"
	"log"
	"os"
)

/*
İf  eğer koşul sağlanmış veya sağlanmamışsa duruma göre kontrol eder.
İf yapısının içerisinde return veya break gibi kontrol ifadeleri kullanılır.

return işlemin sonucunu içerisinde tutar, sen çağırana kadar bir sonuç yazmaz.

* if ve switch genellikle initialize(başlatma) ifadesi olrak kabul edildiği için
genellikle llokal bir değişkeni veya işlemi başalatmak için kullanılır.

* koşul gerçekleşmişse belirtilen işlemi yap,
gerçekleşmemişse belirtilen işlemi yap.
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

	// f, err := os.Open(name)
	//bu kod name dosyasını aç anlamına gelmektedir.

	//f, err : eğer dosya açılırsa hatasız, f e aktarılır.
	// err ise olası bir hatada bu değişkene hata aktarılır.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	//defer : bir fonksiyonun sonlanması ile birlikte belirli işlemi ertelemek için kullanılan kelimedir.
	//fonksiyonun sonuna gelindiği zaman yada bir return ifadesine gelindiğinde otomatik çalıştırır.
	//kardeşim işlem tamamlandıktan sonra şunları yap, veri tabanını kapat.
	//bu sayede kodlar  garanti altına alınır.
	//biz farklı yerlerde işlem açtık, farklı yerlerde işlemleri kapatacağız bu sayede kaynak sızıntısını önleyeceğiz.
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	//d, err := f.Stat()
	//f.stat f teki değerleri d değişkenine aktar.
	//f.Stat() çağırdığımızda, dosya bilgilerini d ye aktarır. os.FİleInfo türünde özellik döndürür.
	fmt.Println("Dosya adı:", info.Name())
	fmt.Println("boyut", info.Size())
	fmt.Println("mod", info.Mode())
	fmt.Println("Bu bir dizimi", info.IsDir())

}
