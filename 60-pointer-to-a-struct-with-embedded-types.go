//show how to use pointers with structs that have embedded types.
package main
import "fmt"

type Address struct {
  city string
}

type User struct {
  name string 
  Address //Embedded Type
}

func updateCity(user *User,newCity string){
  user.city = newCity
}

func main(){
  user := User{name:"Bob", Address: Address{city:"Old City"}}
  updateCity(&user, "New City")
  fmt.Println("Updated user:",user)
}
