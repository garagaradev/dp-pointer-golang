//Write a program where an interface method takes a pointer receiver and demonstrate
//the difference between passing a pointer and a value
package main
import "fmt"

type Printer interface {
  Print()
}

type Document struct {
  content string
}

func (d *Document) Print(){
  fmt.Println(d.content)
}

func main(){
  doc := Document{"Hello World!"}
  var p Printer = &doc
  p.Print()
}
