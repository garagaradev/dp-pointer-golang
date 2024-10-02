//create a struct that contains a pointer to another struct, 
//and write a function to modify it.
package main
import "fmt"

type Engine struct {
  Horsepower int
}

type Car struct {
  Brand string
  Engine *Engine
}

func updateEngine(c *Car, horsepower int){
  c.Engine.Horsepower = horsepower
}

func main(){
  engine := &Engine{Horsepower:100}
  car := &Car{Brand:"Ford", Engine:engine}

  fmt.Println("Before:", car.Engine.Horsepower)
  updateEngine(car, 150)
  fmt.Println("After:", car.Engine.Horsepower)
}
