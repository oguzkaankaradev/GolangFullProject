package main

import "fmt"

//interface : imza : arayüz
// bir veya birden fazla fonksiyonu imzalamak için kullanırız.
//Biz interfaceler ile fonksiyonun kapsamını yani içerininde neler olacağını önceden belirleyip imzalıyoruz.
//biz genel kapsamı yazdıktan sonra functiona implemente ederiz.

//go lang de Interfaceler isimlendirilirken er takısı gelir. read(func) , reader(interface)
//Neden bu kurallar mnemlidir,
//Goda bazı standartlar vardır, bunlar artık yerleşiktir, Read([]byte, error)
//temel veri tipleri ile interface oluşturmamalısın, string -> stringer bu isimde oluşturma

//veri tabanı Connect, Disconnet olur
type Connector interface {
	Connect() error
	Disconnect() error
}

//Veri tabanından veri okuruz
type Reader interface {
	Read(query string) ([]byte, error)
}

//Veri tabanına veri yazarız
type Writer interface {
	Write(data []byte) error
}

//veri tabanının sağlığını test ederiz
type HealtChecker interface {
	IsConnected() bool
}

//Bunları tek çatı  altında toplanması, bunun üzerinden ulaşacağız.
type DataStore interface {
	Connector
	Reader
	Writer
	HealtChecker
}

//interface, veri tabanı işlemlerini tanımlar ama nasıl uygulayacağını bilemez.
//Nasıl uygulayacağı sorusunun cevabını DataBase gibi structlar verir.
type DataBase struct {
	connected bool
}

//
func (d *DataBase) Connect() error {
	d.connected = true
	fmt.Println("Veri tabanına bağlantı başarı ile kuruldu")
	return nil
}

func (d *DataBase) Disconnect() error {
	d.connected = false
	fmt.Println("Veri tabanı bağlantısı kurulamadı")
	return nil
}

func (d *DataBase) Read(query string) ([]byte, error) {
	if !d.connected {
		return nil, fmt.Errorf("Veri tabanı bağlantısı kurulamadı")
	}
	fmt.Println("Veri tabanı okundu")
	return []byte("veri"), nil
}

func (d *DataBase) Write(data []byte) error {
	if !d.connected {
		return fmt.Errorf("Veri tabanı bağlantısı kurulamadı")
	}
	fmt.Println("Veri kaydedildi", string(data))
	return nil
}

func (d *DataBase) IsConnected() bool {
	return d.connected
}

//
type Animal interface {
	speak() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) speak() string {
	return "hav hav"
}

func (c Cat) speak() string {
	return "Miyav"
}

func MakeSpeak(a Animal) {
	a.speak()
}

type BlogService interface {
	DeleteBlogPost() error
	BlogPostById() error
	AddBlog() error
	UpdeteBlog() error
	ParseTag() error
	AllCommets() error
}

//blog interface sini uygular
type BlogMake struct {
}

func (b BlogMake) DeleteBlogPost() error {
	//....
	fmt.Println("Bloğunuz başarı ile silindi")
	return nil
}

func (b BlogMake) AddBlog() error {
	//...
	fmt.Println("Blok yazınız eklendi")
	return nil
}

func (b BlogMake) AllCommets() error {
	//...
	fmt.Println("Blog yorumları yüklendi")
	return nil
}

func (b BlogMake) BlogPostById() error {
	fmt.Println("Aradığınız blog bulundu")
	return nil
}

func (b BlogMake) ParseTag() error {
	fmt.Println("Taglarınız pars edildi")
	return nil
}

func (b BlogMake) UpdeteBlog() error {
	fmt.Println("Blok yazınız başarı ile güncellendi")
	return nil
}

func main() {
	var store DataStore = &DataBase{}

	store.Connect()

	if store.IsConnected() {
		store.Write([]byte("isim oraz"))
		result, _ := store.Read("SELECTED * FROM users")
		fmt.Println("Okunan veri", string(result))
	}

	var a Animal

	a = Dog{}
	MakeSpeak(a)

	a = Cat{}
	MakeSpeak(a)

	var blog BlogService = BlogMake{}

	blog.AddBlog()
	blog.DeleteBlogPost()
	blog.AllCommets()
	blog.BlogPostById()
	blog.ParseTag()
	blog.UpdeteBlog()

}
