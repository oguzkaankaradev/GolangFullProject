package semicolons

import (
	"fmt"
	"time"
)

//Noktalı virgül ekleme kurallarının bir sonucu olarak, bir kontrol yapısının (if, for, switch veya select) açılış süslemesini bir sonraki satıra koyamazsınız. Eğer yaparsanız, süslemenin öncesine bir noktalı virgül eklenir ki bu istenmeyen etkilere yol açabilir. Bu şekilde yazmalısınız:

////Önemli kesin kullan.
///func() { for {dst <- <- src}}()

// go : yeni bir go runtime başlatmak için kullanılır.
//func() {} : bu bir anonim fonksiyondur
//for{} sonsuz bir for döngüsü oluşturur.
//<-src: src bir kanaldan veri almak için
//dst : veriyi başka bir kanala görndermek için

//Bu yapı go runtimında bir sonsuz döngü oluşturur, src kanalından verileri alır dst adlı başka bir kanala gönderir.

func main() {
	//make : Dinamik olarak bellek ayırmak için kullanılır.
	src := make(chan int)
	dst := make(chan int)

	//goruntime ile veriyi aktarıyoruz.
	go func() {
		for {
			dst <- <-src // src kanalından veriyi al dst ye gönder.

		}
	}()

	//src kanalına veri gönderir
	go func() {
		for i := 0; i < 5; i++ {
			src <- i
			time.Sleep(time.Second) // Her bir gönderim arasında bekle
		}
		close(src) // src kanallını kapatıyoruz.

	}()

	//dst kanalından verileri okuma
	for val := range dst {
		fmt.Println(val)
	}
}

//Not kanaldan kanala veri aktarırken, iki tane <- <- kullanılır.
//Değişkenden kanala veri aktarırken bir tane <- kullanılır.
