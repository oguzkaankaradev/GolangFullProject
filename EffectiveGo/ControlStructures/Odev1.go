package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// bir kafede, müşterilerin adını, soyadını yaşı ve bilgisayarına
// erişerek tarihi yazın. Müşteriler txt dosyasına sırayla kaydolmasını sağla.
// ve müşerileri eğer 18 de büyük ve eşitse giriş yapan müşteriler klasörüne,
// 18 den küçükse giriş yapamayan müşteriler klasörüne kaydet.

func Odev1() {
	var firstName string
	var lastName string
	var fullName string
	fullName = firstName + " " + lastName
	var age int                  //age := int
	var currentTime = time.Now() //güncel zaman.
	var fileName string
	var folder string

	fmt.Print("Adınızı giriniz:")
	fmt.Scanln(&firstName)

	fmt.Print("SoyAdınızı giriniz:")
	fmt.Scanln(&lastName)

	fmt.Print("Yaşınızı giriniz:")
	fmt.Scanln(&age)

	if age < 18 {

		println("girme")
		return

	} else {

		formatedTime := currentTime.Format("2006-01-02 15:04:05")

		fileName = "Musteriler.txt"

		folder = "Customer"
		//MkdirAll : Klasör oluşturma kodu.
		// folder oluşturulacak klasör
		//os.ModePerm : klasörlerin erişim izinleri. Tüm izinler verilir bu kod ile(okuma yazma silme çalıştırma)
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			fmt.Println("klasör oluşturulamadı", err)
		}
		if age >= 18 {
			folder = "giriş yapan müşteriler kodu"
		} else {
			folder = "giriş yapamayan müşteriler"
		}

		//filepath.Join : farklı işletim sistemlerinde(win, mac) dosya yollarının nasıl birleştirileceğini otomatik belirler. /
		filePath := filepath.Join(folder, fileName)

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
		formatedString := fmt.Sprintf("fullname: %s, yaş: %d, Tarih: %s", fullName, age, formatedTime)
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

		fmt.Println("helal olsun.qw")
	}

}
