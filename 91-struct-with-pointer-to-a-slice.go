//create a struct that contains a pointer to a slice and
//write a function to add elements to the slice.
package main
import "fmt"

type Group struct {
  Members *[]string
}

func addMember(g *Group, member string){
  *g.Members = append(*g.Members, member)
}

func main(){
  members := []string{"Alice"}
  group := &Group{Members:&members}

  addMember(group, "Bob")
  fmt.Println("Group Members:", *group.Members)

}
