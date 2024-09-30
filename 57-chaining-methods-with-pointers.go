//demonstrate method chaining using structs
package main
import "fmt"

type Builder struct {
  name string
}

func (b *Builder) SetName(name string) *Builder {
  b.name = name
  return b
}

func (b *Builder) Build() string {
  return "Building " + b.name
}

func main(){
  builder := &Builder{}
  result := builder.SetName("House").Build()
  fmt.Println(result)
}
