package python

import (
	"fmt"

	"github.com/elnaterator/bootstrapper/pkg/core"
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
	fmt.Println("\nBootstrapping a python project...")
}
