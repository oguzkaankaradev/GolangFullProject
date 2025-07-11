//Go formatlama
//Bir programda kodun herkesin aynı şekilde görmesidir. Herkesin aynı kurallara syntax a uymasıdır.
//gofmt : Kodlarınızı otomatik düzenleyen bir araçtır.


type T struct {
	name  string
	value int
}

//go da maxsimum satır uzunluğu diye bişey yoktur.

type T struct {
	name  string            value int  
}

//Goda parantez kullanılmaz

if(x + y > 10 ) {......} //yanlış formattır. c# gibi dillerde parantez varken goda yoktur. 
if x + y > 10 { ...... }