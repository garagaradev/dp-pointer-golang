//implement a deep copy function for a struct with pointer fields
package main
import "fmt"

type Address struct {
  city *string
}

type User struct {
  name string
  address *Address
}

func deepCopy(user *User) User {
  newAddress := &Address{city: user.address.city}
  return User{name: user.name, address: newAddress}
}

func main(){
  city := "New York"
  original := User{name: "Alice", address: &Address{city: &city}}
  copied := deepCopy(&original)

  fmt.Println("Original:",*original.address.city)
  fmt.Println("Copied:",*copied.address.city)

  city = "Los Angeles"
  fmt.Println("After modifying city...")
  fmt.Println("Original:",*original.address.city)
  fmt.Println("Copied:",*copied.address.city)
}
