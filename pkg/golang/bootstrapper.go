package golang

import (
	"strings"

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
	version    string
	module     string
	command    string
}

func (b *Bootstrapper) CollectAdditionalOptions() {

	b.fetchGoVersions()
	version := input.Choice("Go version?", b.goVersions)
	b.version = strings.Replace(version, "go", "", 1) // we just want the number

	module := input.Text("Module name?")
	b.module = module
	parts := strings.Split(module, "/")
	b.command = parts[len(parts)-1]
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
			makefile(b.command),
			readme(b.command),
			gitignore(),
			goVersion(b.version),
			goMod(b.version, b.module),
		},
		Directories: []core.Directory{
			{
				Name: "cmd",
				Directories: []core.Directory{
					{
						Name:  b.command,
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
	return core.Traverse(&b.project.RootDir, core.WriteFileHandler, nil)
}
