package manin

import (
	"fmt"
	"os"
)

//Reassignment - Redeclaration - Redirection
// := bu işaret hem değişkeni tanımlamak için hemde değişkeni yeniden kullanmak için kullanılır.

func main() {
	f, err := os.Open("test.txt") //f ve err tanımlandı
	if err != nil {
		return err
	}

	d, err := f.Stat() //d yeni tanımlandı err tekrar kullanıldı
	if err != nil {
		f.Close()
		return err
	}

	//
	x := 1
	y := 2

	x, z := 2, 4 //x tekrar atandı.

	fmt.Println(x, y, z)
}
