package python

import (
	"github.com/elnaterator/bootstrapper/pkg/core"
)

func NewBootstrapper(project *core.Project) core.Bootstrapper {
	return &Bootstrapper{
		project: project,
	}
}

type Bootstrapper struct {
	project *core.Project
}

func (b *Bootstrapper) CollectAdditionalOptions() {
}

func (b *Bootstrapper) BuildModel() *core.Directory {
	return nil
}

func (b *Bootstrapper) Bootstrap() {
}
