package main

import "fmt"

func main() {

	m := map[string]*int{}

	m["a"] = nil
	m["b"] = new(int)
	*m["b"] = 1

	fmt.Println(m)

}
