//illustrate how pointers to interface work
package main
import "fmt"

type Printer interface {
  Print()
}

type Document struct {
  title string
}

func (d *Document) Print(){
  fmt.Println("Printing document:",d.title)
}

func main(){
  doc := &Document{title:"Go Interface"}
  var p Printer = doc
  p.Print()
}
