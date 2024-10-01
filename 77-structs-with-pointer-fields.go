//create a struct that contains a pointer field and demonstrates how to modify it.
package main
import "fmt"

type Config struct {
  Setting *string
}

func main(){
  setting := "Enabled"
  config := Config{Setting: &setting}
  fmt.Println("Before:",*config.Setting)
  newSetting := "Disabled"
  config.Setting = &newSetting
  fmt.Println("After:",*config.Setting)
}
