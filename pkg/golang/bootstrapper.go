package golang

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
	b.project.RootDir = core.Directory{
		Name: b.project.RootDirName,
		Files: []core.File{
			{Name: "Makefile", Content: ""},
			{Name: "README.md", Content: ""},
		},
		Directories: []core.Directory{
			{Name: "cmd"},
			{Name: "pkg"},
		},
	}
	return &b.project.RootDir
}

func (b *Bootstrapper) Bootstrap() {
}
