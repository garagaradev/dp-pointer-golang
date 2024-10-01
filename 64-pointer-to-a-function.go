//create a function that takes a pointer to another function as an argument and calls it.
package main
import "fmt"

func greet(name string){
  fmt.Println("Hello ",name)
}

func callFunction(fn *func(string), name string){
  (*fn)(name)
}

func main(){
  greetFunc := greet
  callFunction(&greetFunc, "World!")
}
