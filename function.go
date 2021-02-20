//Example -6:

package main

import "fmt"

func update(a *int , t * string) {
	*a = *a + 5
	*t = *t + "Doe"
	return 
}

func main() {
	number := 10
	name :="Sourav"
	update(&number , &name)
	fmt.Println(number, name)
}
