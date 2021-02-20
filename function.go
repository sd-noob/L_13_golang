//Example -4:(Name return value)

package main

import "fmt"

func add(x , y int) (r int){
r = x+y
return r
}

func main(){
x := add( 10, 30)
fmt.Println(x)

}