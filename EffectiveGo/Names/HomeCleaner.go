package names

import "fmt"

func main() {

	fmt.Println("Gardaş senin görevin evi temizlemek.")
}

//Get set example

type Person struct {
	name     string // private field
	age      int    // private field
	identity string
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(newName string) {
	newName = "salih"
	p.name = newName

}
