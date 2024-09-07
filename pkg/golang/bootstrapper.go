package golang

import (
	"fmt"

	"github.com/elnaterator/bootstrapper/pkg/core"
	"github.com/elnaterator/bootstrapper/pkg/input"
)

type Bootstrapper struct {
	project *core.Project
}

func NewBootstrapper(project *core.Project) *Bootstrapper {
	return &Bootstrapper{
		project: project,
	}
}

func (b *Bootstrapper) Bootstrap() {
	fmt.Println("\nBootstrapping a golang project...")

	goVersion := input.Choice("golang version?", []string{"1.23"})
	fmt.Println(goVersion)

}
