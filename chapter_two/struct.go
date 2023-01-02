package chapter_two
import "fmt"
type Person struct {
	Name string
	Age int
}

func newPerson(s string) *Person {

	p :=Person{Name:s}

	p.Age = 42
	return &p

}
func PrintSome() {
  fmt.Println("chapter_two")

}
func Create(n string) *Person {
	p :=Person{Name:n}
	p.Age =45
	return &p
}
