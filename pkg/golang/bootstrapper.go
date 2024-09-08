package golang

import (
	"github.com/elnaterator/bootstrapper/pkg/core"
	"github.com/elnaterator/bootstrapper/pkg/input"
)

func NewBootstrapper(project *core.Project) core.Bootstrapper {
	return &Bootstrapper{
		project: project,
	}
}

type Bootstrapper struct {
	project    *core.Project
	goVersions []string
}

func (b *Bootstrapper) CollectAdditionalOptions() {
	b.fetchGoVersions()
	input.Choice("Go version?", b.goVersions)
}

func (b *Bootstrapper) fetchGoVersions() {
	resp := core.GetJson("https://golang.org/dl/?mode=json&include=all").([]interface{})
	for _, v := range resp {
		version := v.(map[string]interface{})["version"]
		b.goVersions = append(b.goVersions, version.(string))
	}
}

func (b *Bootstrapper) BuildModel() *core.Directory {
	b.project.RootDir = core.Directory{
		Name: b.project.RootDirName,
		Files: []core.File{
			makefile(),
			readme(),
			gitignore(),
		},
		Directories: []core.Directory{
			{
				Name: "cmd",
				Directories: []core.Directory{
					{
						Name:  "app",
						Files: []core.File{mainGo()},
					},
				},
			},
			{Name: "pkg"},
		},
	}
	return &b.project.RootDir
}

func (b *Bootstrapper) Bootstrap() error {
	return nil
}
