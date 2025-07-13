package main

//Getter ne işe yarar : Veriyi okumak için kullanılan fonksiyondur.
//Setter ne işe yarar : Veriyi yazmak için kullanılır, güncelleme.

//Getter Setter
//goda otomatik getter ve setter mekanizması yoktur java ve dotnet gibi dillerde vardır.
//ihtiyaç dıyduğumuzda manuel olarak yazacağız.

func main() {

}

//yazarken goya uygun syntax ile yazacağız.

//yanlış isimlendirme, bu zaten bir get funcsion u dur. tekrardan get yazmana gerekyok.
// func (p *Person) GetName() string {
// 	return p.Person
// }

// //doğru yazılış
// func (p *Person) Name() string {
// 	return p.Person
// }

type User struct {
	name string
}

func (u *User) Name() string {
	return u.name
}

//Gette başına get kullanmazken sette set kullanırız.
//zamanla set içerisinde çeşitli güvenlik önlemleri ve loglama ve hata yönetimi gibi işlemler yapıcaz.
func (u *User) SetName(n string) {
	if n != "" {
		u.name = n
	}
}

//////////////////////////////////////

type Account struct {
	balance float64
}

//bakiyeyi okunmak için get
func (account *Account) Balance() float64 {
	return account.balance
}

//bakiye ekleyelim içine set ile, ama sadece pozitif değer ekleyelim.
func (account *Account) SetBalance(amount float64) {
	if amount >= 0 {
		account.balance = amount
	}
}
