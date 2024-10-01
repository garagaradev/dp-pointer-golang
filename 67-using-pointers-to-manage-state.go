//create a program that uses pointers to manage the state of a game character
package main
import "fmt"

type Character struct {
  name string
  health int
}

func heal(c *Character, amount int){
    c.health += amount
}

func main(){
  hero := &Character{name:"Warrior",health:50}
  fmt.Println("Before healing:", hero.health)
  heal(hero,100)
  fmt.Println("After healing:", hero.health)
}
