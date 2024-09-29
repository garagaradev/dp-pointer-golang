// write a program that demonstrates passing a pointer to a function as a callback
package main
import "fmt"

func process(value int, callback func(int)){
  callback(value)
}

func printDouble(n int){
    fmt.Println(2 * n)
}

func main(){
  process(15, printDouble)
}
