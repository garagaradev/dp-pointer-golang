//create a function that modifies a struct stored in a slice via a pointer
package main
import "fmt"

type Employee struct {
  Name string
  Age int
}

func updateEmployee(emp *Employee, name string, age int){
  emp.Name = name
  emp.Age = age 
}

func main(){
  employees := []*Employee{{Name:"Jack",Age:38}}
  updateEmployee(employees[0], "Dorsey",39)
  fmt.Println("Updated Employee:", employees[0])
}
