//create a function that accepts a pointer to an interface and modifies it.
package main
import "fmt"

type Animal interface {
  Speak() string
}

type Dog struct {}

func (d *Dog) Speak() string {
  return "Woof!"
}

func changeAnimal(a *Animal){
  dog := &Dog{}
  *a = dog
}

func main(){
  var animal Animal
  changeAnimal(&animal)
  fmt.Println(animal.Speak())
}
