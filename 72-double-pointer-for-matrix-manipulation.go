//create a function that manipulates a 2D slice (matrix) using a double pointer
package main
import "fmt"

func setMatrixValue(m *[][]int, row, col, value int){
  (*m)[row][col] = value
}

func main(){
  matrix := [][]int{{1,2},{3,4}}
  fmt.Println("Before:",matrix)
  setMatrixValue(&matrix, 1, 0, 99 )
  fmt.Println("After:", matrix)
}
