//Write a function that modifies a string using a pointer
package main
import "fmt"

func modifyString(s *string){
  *s = "Modified String"
}

func main(){
  str := "Original String"
  modifyString(&str)
  fmt.Println("Modified String:", str)
}
