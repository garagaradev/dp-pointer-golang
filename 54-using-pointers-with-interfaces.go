//implement a function that accepts an interface
//and modifies its underlying value using a pointer.
package main
import "fmt"

type Stringer interface {
  String() string
}

type Person struct {
  name string
}

func (p *Person) String() string {
  return p.name
}

func modifyStringer(s Stringer){
  if p, ok := s.(*Person); ok {
    p.name = "Modified Name"
  }
}

func main(){
  person := &Person{name:"Original Name"}
  modifyStringer(person)
  fmt.Println("Modified Person Name:",person.String())
}

