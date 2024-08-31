package greet

import "fmt"

type Greeter struct {
	Name string
}

func (h *Greeter) BuildGreeting() string {
	return fmt.Sprintf("Hello %s, how are you?", h.Name)
}
