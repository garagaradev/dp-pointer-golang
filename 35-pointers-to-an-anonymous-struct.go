//write a program that demonstrates the use of pointers to anonymous struct 
package main
import "fmt"

func main(){
  student := &struct {
    name string
    grade int
  }{
    name: "Alice",
    grade: 95,
  }

  student.grade = 98
  fmt.Printf("%s's grade: %d\n",student.name,student.grade)
}
