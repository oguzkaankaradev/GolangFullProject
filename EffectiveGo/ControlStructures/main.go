package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	// var nargileKafeYas = 18
	// var CustomerAge int
	// fmt.Println("Lütfen yaşınızı giriniz")
	// fmt.Scan(&CustomerAge)

	// if CustomerAge >= nargileKafeYas {
	// 	fmt.Println("Hoşgeldiniz")
	// } else {
	// 	fmt.Println("velinizle geliniz.")
	// }

	var firstName string
	var lastName string
	var age int                  //age := int
	var currentTime = time.Now() //güncel zaman.
	var fileNameDown string
	var fileNameUp string
	var folder string

	fmt.Print("Adınızı giriniz:")
	fmt.Scanln(&firstName)

	fmt.Print("SoyAdınızı giriniz:")
	fmt.Scanln(&lastName)

	fmt.Print("Yaşınızı giriniz:")
	fmt.Scanln(&age)

	formatedTime := currentTime.Format("2006-01-02 15:04:05")

	fileNameUp = "Musteriler.txt"
	fileNameDown = "kucukler.txt"

	folder = "Customer"

	//MkdirAll : Klasör oluşturma kodu.
	// folder oluşturulacak klasör
	//os.ModePerm : klasörlerin erişim izinleri. Tüm izinler verilir bu kod ile(okuma yazma silme çalıştırma)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		fmt.Println("klasör oluşturulamadı", err)
	}

	if age >= 18 {
		//filepath.Join : farklı işletim sistemlerinde(win, mac) dosya yollarının nasıl birleştirileceğini otomatik belirler. /
		filePath := filepath.Join(folder, fileNameUp)

		//filePath : oluşturulacak olan dosyanın yolu.
		//os.O_APPEND : bir dosya var, içerisinde veri var, o verinin üzerine yazmak yerine, altına yeni veriyi yazar.
		//os.O_CREATE : belirtilen dosya yoksa oluşturur.
		//os.O_WRONLY : Dosyanın yalnızca yazma modunun açılacağını belirtir, dosya üzerinde okuma yapılmayacağını belirtir.
		//6 (owner okuma ve yazma iznine sahip),
		//4 (group okuma iznine sahip)
		//4 (other users okume iznine sahip)
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("dosya açılmadı:", err)
			return
		}

		defer file.Close()

		//bufio : paketin ismi. bu paket dosya yazmadan önce bellekte toplanmasını sağlar.
		//NewWriter : yeni dosyayı bellekte tutuyor.
		// amaç performansı ertırmak. verimliliği artırmak.
		writer := bufio.NewWriter(file)
		formatedString := fmt.Sprintf("name: %s, lastname %s, yaş: %d, Tarih: %s", firstName, lastName, age, formatedTime)
		_, err = writer.WriteString(formatedString)

		if err != nil {
			fmt.Println("Yazma hatası", err)
			return
		}

		// writer.Flush() : yazma işlemini daha verimli hale getirmek için depoda toplar.
		writer.Flush()
		if err != nil {
			fmt.Println("Flush hatası", err)
		}

		println("Helal olsun")

	} else {
		//filepath.Join : farklı işletim sistemlerinde(win, mac) dosya yollarının nasıl birleştirileceğini otomatik belirler. /
		filePath := filepath.Join(folder, fileNameDown)

		//filePath : oluşturulacak olan dosyanın yolu.
		//os.O_APPEND : bir dosya var, içerisinde veri var, o verinin üzerine yazmak yerine, altına yeni veriyi yazar.
		//os.O_CREATE : belirtilen dosya yoksa oluşturur.
		//os.O_WRONLY : Dosyanın yalnızca yazma modunun açılacağını belirtir, dosya üzerinde okuma yapılmayacağını belirtir.
		//6 (owner okuma ve yazma iznine sahip),
		//4 (group okuma iznine sahip)
		//4 (other users okume iznine sahip)
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("dosya açılmadı:", err)
			return
		}

		defer file.Close()

		//bufio : paketin ismi. bu paket dosya yazmadan önce bellekte toplanmasını sağlar.
		//NewWriter : yeni dosyayı bellekte tutuyor.
		// amaç performansı ertırmak. verimliliği artırmak.
		writer := bufio.NewWriter(file)
		formatedString := fmt.Sprintf("name: %s, lastname %s, yaş: %d, Tarih: %s", firstName, lastName, age, formatedTime)
		_, err = writer.WriteString(formatedString)

		if err != nil {
			fmt.Println("Yazma hatası", err)
			return
		}

		// writer.Flush() : yazma işlemini daha verimli hale getirmek için depoda toplar.
		writer.Flush()
		if err != nil {
			fmt.Println("Flush hatası", err)
		}

		fmt.Println("Girme")

	}

}
