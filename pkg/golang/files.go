package golang

import (
	_ "embed"
	"strings"

	"github.com/elnaterator/bootstrapper/pkg/core"
)

//go:embed Makefile
var Makefile string

func makefile(command string) core.File {
	content := strings.Replace(Makefile, "command", command, -1)
	return core.BuildFile("Makefile", content)
}

//go:embed README.md
var Readme string

func readme(command string) core.File {
	title := strings.ToTitle(command)
	content := strings.Replace(Readme, "Title", title, 1)
	content = strings.Replace(content, "command", command, -1)
	return core.BuildFile("README.md", content)
}

//go:embed main.go.txt
var MainGo string

func mainGo() core.File {
	return core.BuildFile("main.go", MainGo)
}

func gitignore() core.File {
	return core.BuildFile(".gitignore", ".idea\n")
}

func goVersion(version string) core.File {
	return core.BuildFile(".go-version", version)
}

func goMod(version string, module string) core.File {
	return core.BuildFile("go.mod", "module "+module+"\n\ngo "+version)
}
