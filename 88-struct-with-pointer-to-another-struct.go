//create a struct that contains a pointer to another struct and modify it.
package main
import "fmt"

type Address struct {
  City string
  State string
}

type Person struct {
  Name string
  Address *Address
}

func updateAddress(p *Person, city, state string){
    p.Address.City = city
    p.Address.State = state
}

func main(){
  addr := &Address{City:"New York",State:"NY"}
  person := &Person{Name:"Alice", Address:addr}

  fmt.Println("Before:", person.Address)
  updateAddress(person, "Los Angeles","CA")
  fmt.Println("After:", person.Address)
}


