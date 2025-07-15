package main

import "fmt"

//Semicolons
//Noktalı virgül  ekleme kuralı
//Bir satırdaki son token(birim) şunlardansa
//Bir tanımlayıcı(int, float, x,myvariable),
//Sabit bir değer ise("merhaba",123,thrue, vb)
//şu özel anahtar sözcüklerden biri ise(break, continue, fallthrough, return, ++,--,), })
//ve bu tokendan hemen sonra yeni bir satıra geçilmişse, go otomatik olarak satırın sonuna ; ekler.

func main() { //bu parantezi alt satıra koamazsın,
	// fmt.Println("merhaba")
	// return

	//For döngüsündeki ifadeleri birbirinden ayırmak için ; kullanılır.
	for i := 0; i < 500; i++ {

	}

	//Aynı satırda birden fazla ifade varsa, bu ifadeleri ayırmak için kullanılmaktadır.
	x := 10
	fmt.Println(x)
}

//Yukarıdaki örneklerdede anlattığımız gibi ; ifadesi go lang de çoğu yerde kullanılmamaktadır, lexer kendisi otomatik koyar.
//ama parantezleri satır sonuna yazamazsın beklenmedik hatalar ile karşılaşırsın.
//{ yazım kuralı gereği daima aynı satırda olmalıdır.
