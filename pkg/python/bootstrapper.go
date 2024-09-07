package python

import "fmt"

type Bootstrapper struct{}

func NewBootstrapper() *Bootstrapper {
	return &Bootstrapper{}
}

func (b *Bootstrapper) Bootstrap() {
	fmt.Println("\nBootstrapping a python project...")
}
