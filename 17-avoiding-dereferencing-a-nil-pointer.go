//write a program that handles a nil pointer properly
package main
import "fmt"

func safeDereference(p *int){
  if p == nil {
    fmt.Println("Pointer is nil, cannot dereference")
  }else {
    fmt.Println("Value:",*p)
  }
}

func main(){
  var p *int
  safeDereference(p)

  x := 10
  p = &x

  safeDereference(p)

}
