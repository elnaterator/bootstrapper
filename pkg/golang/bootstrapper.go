package golang

import (
	"fmt"

	"github.com/elnaterator/bootstrapper/pkg/options"
)

type Bootstrapper struct{}

func NewBootstrapper() *Bootstrapper {
	return &Bootstrapper{}
}

func (b *Bootstrapper) Bootstrap() {
	fmt.Println("\nBootstrapping a golang project...")

	goVersion := options.GetOptions("golang version?", []string{"1.23"})
	fmt.Println(goVersion)

}
