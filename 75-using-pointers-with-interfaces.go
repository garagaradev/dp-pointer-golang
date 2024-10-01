// create a program that uses a pointer with an interface type
package main
import "fmt"

type Stringer interface {
  String() string
}

type Person struct {
  Name string
}

func (p *Person) String() string {
  return p.Name
}

func printName(s Stringer){
  fmt.Println(s.String())
}

func main(){
  person := &Person{Name:"Alice"}
  printName(person)
}
