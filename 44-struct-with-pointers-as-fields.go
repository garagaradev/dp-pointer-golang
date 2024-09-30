//create a struct that has pointers as fields and demonstrate their usage
package main
import "fmt"

type Person struct {
  name *string
  age *int
}

func main(){
  name := "Jack"
  age := 30
  p := Person{name: &name, age:&age}
  fmt.Println("Name:",*p.name)
  fmt.Println("Age:",*p.age)
}
