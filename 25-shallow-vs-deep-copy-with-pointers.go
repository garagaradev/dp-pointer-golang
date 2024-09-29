//illustrate the difference between shallow and deep copying using pointers in structs
package main
import "fmt"

type Person struct {
  name string
  friend *Person
}

func main(){
  jack := &Person{name: "Jack"}
  mike := &Person{name:"Mike", friend:jack}

  shallowCopy := mike
  deepCopy := *mike
  deepCopy.friend = &Person{name:"New Friend"}

  fmt.Println(shallowCopy.friend.name)
  fmt.Println(deepCopy.friend.name)
}
