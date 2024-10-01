//Write a function that returns a pointer to a newly allocated struct
package main
import "fmt"

type Car struct {
  Make string
  Model string
}

func newCar(make, model string) *Car {
  return &Car{Make:make, Model:model}
}

func main(){
  car := newCar("Toyota", "Camry")
  fmt.Printf("Car: %+v\n", car)
}
