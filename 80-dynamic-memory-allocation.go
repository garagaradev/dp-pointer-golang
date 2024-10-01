//Write a program that allocates memory from a struct dynamically using pointers
package main
import "fmt"

type Student struct {
  Name string
  Grade int
}

func main(){
  student := new(Student)
  student.Name = "Jack Dorsey"
  student.Grade = 90
  fmt.Printf("Student:%+v\n",student)
}
