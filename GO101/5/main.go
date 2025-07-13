package main //yazdığımız kodu package ile paket içerisine alırız. yeni bir package oluşturduk.
//  artık erişebiliriz import ile.
//Package Name
//Biz kodlarınmızı yazarız paketleriz, sonradan kullanırız.
// Her go projesi mutlaka paketlendirilir
//import kelimesi ile daha önce yazdığımız packageye veya daha önce yazılan packageye erişiriz.
//Projemize dahil ederiz.
//Paket isimleri anlamlı, kısa ve teke kelime olmalıdır.
//biz
import "bytes"

func main() {
	var b bytes.Buffer

	b.WriteString("hello")

}

//paket isimlendirir ken _ kullanmayın
//Büyük harf kullanmayın
//çok uzun isimlerden kaçının
//tekrardan kaçının öreneğin bufio.bufread bu yanlıştır. Doğrusu bufio.read
