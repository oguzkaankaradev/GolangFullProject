//Names : İsimlendirme kurallarını işleyeceğiz. 
//Go langda isimlendirme kuralları vardır, ve bu isimlendirme kuralları teknik olarak anlam taşımaktadır. 

//En önemli iki kural vardır
//1. private isimlendirme = Küçük harfle başlaması : ilk harf küçük harfle başlarsa, private olur. Yalnızca aynı paket içerisinde kullanabilirsin. 
var internalValue = 10
func doSomething(){
	...
}

//2. Public isimlendirme = Büyük harfle başlaması : İlk harf büyük harfle başlar ise publictir, yani diğer paketlerdende erişilebilmektedir. 
var ExportValur = 10
func DoSomething(){
	...
}

//Kural : ilk harf küçük ise dışa kapalıdır, büyükse dışa açıktır. 


//iyi isimlendirme neden önemlidir
//Kodunuz okunur ve anlaşılır yapar
//Takım arkadaşlarınız rahatlıkla geliiştirir
//Kodun sürekliliğini sağlar. Kodu review etmen kolay olur. 


//İyi isimlendirme nasıl olur 

//1. Açıklayıcı olun kısa net. 
func Sum(a,b int) int
func AddNumber(a, b int) int //dikkat et

//2. Türlerle uyumlu ol
type User struct{
	Name string
	Age int
}

func NewUser(name string, age int) User{
	return User{Name:name, Age:age}
}

//3. Getter ler ile başlayan isimlendirmenin başına Get koymanıza gerek yok. 
func GetName() //buna gerek yok
func Name()

//Kısaltmaları Dikkatli Kullanın
func XmlToJson(){...} //bu yanlış
func XMLToJSON(){...}