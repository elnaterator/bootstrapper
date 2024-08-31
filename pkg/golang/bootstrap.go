package golang

import "fmt"

type Bootstrapper struct {
	Command string
}

func (b *Bootstrapper) GatherOptions() {

	fmt.Println("Type of golang project?")
	// check out https://github.com/charmbracelet/bubbletea/blob/master/examples/list-simple/main.go for this

}
