package main

import (
	"log"
	"os"
)

//if
//go dilinde if yapısı c ve java gibi benzer görünsede önemli farklılıklar vardır
//Gonun if yapısı hem sade hemde okunabilirliği artıracak şekilde yazılmıştır

func main() {

	//temel if yapısı
	//parantez() kullanılmaz if (x > 0)
	//{} süslü parantez zorunludur, tek satırlık kod yazsanız bilie zorunludur

	x := 0
	if x > 0 {
		//return ile geriye herhangi birşey döndür ve hatayı gider
		y := 100
		return y
	}

	//Initialization statement(başlatma ifadesi) değişken yani var kullanma
	//if ve switch  başında değişken tanımlayabilirzi.
	//değişken yalnızca if ve altındaki blokta geçerlidir.
	if err := file.chmod(0664); err != nil {
		log.Print(err)
		return err
	} //bu err sadece bu if bloğundadır.

	//else kullanımı
	//goda mümkün olduğunca  else kullanılmaz,
	//çünkü return ile hata kontrolü yaparız, döngüden çıkarız
	//kod akışı daha temiz ve okunabilir oluyor

	//1.
	if err := file.chmod(0664); err != nil {
		return err
	} else {
		codeUsing(f)
	}

	//2
	if err := file.chmod(0664); err != nil {
		return err
	}
	codeUsing(f)

	//hata kontrölü(error handling)
	//goda hata kontrolü yapılırken genelde bir dizi if yazılır ve her hatada return edilir
	//bu yaklaşıma happy path(başarılı yol) denir.

	f, err := os.Open(name)
	if err != nil {
		return err
	}

	d, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}

	codeUsing(f, d)
}
