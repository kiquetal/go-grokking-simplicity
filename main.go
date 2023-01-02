package main
import(
"fmt"
"simplicity/chapter_one"
c "simplicity/chapter_two"
)
func main() {
   fmt.Println("from main")
   chapter_one.PrintHello()
   fmt.Println("check deps")
   c.PrintSome()
   c.Create("Juan")
   p2 := &c.Person{Name:"Kiquetal",Age:21}
   p3 := p2
   p3.Age=34
   println(p2.Age)


}
