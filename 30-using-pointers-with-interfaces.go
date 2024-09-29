// demonstrate how pointers work with interfaces by modifying an object through an interface 
package main
import "fmt"

type Resizeable interface {
  Resize(factor int)
}

type Box struct {
  size int
}

func (b *Box) Resize(factor int){
  b.size *= factor
} 

func main(){
  box := &Box{size:10}
  var r Resizeable = box

  r.Resize(999)
  fmt.Println(box.size)

}
