//illustrate how pass by value and pass by reference affect structs
package main
import "fmt"

type Rect struct {
  width, height int
}

func resizeByValue(r Rect){
  r.width *= 2
  r.height *= 2
}

func resizeByReference(r *Rect){
  r.width *= 2
  r.height *= 2
}

func main(){
  rect := Rect{10,20}

  resizeByValue(rect)
  fmt.Println("After resizeByValue:",rect)

  resizeByReference(&rect)
  fmt.Println("After resizeByReference:",rect)
}
