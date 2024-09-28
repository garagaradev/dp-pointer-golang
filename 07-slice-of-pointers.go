// write a program that creates a slice of integer pointers, 
// modifies the values through the slice, and prints the results
package main
import "fmt"

func main(){
  x, y, z := 1,2,3
  fmt.Println("Before:", x, y, z)
  slice:= []*int{&x,&y,&z}
  
  for _, p := range slice {
    *p = *p * 2
  }

  fmt.Println("After:",x, y, z)

}
