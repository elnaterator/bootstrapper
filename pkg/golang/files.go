package golang

import (
	_ "embed"

	"github.com/elnaterator/bootstrapper/pkg/core"
)

//go:embed Makefile
var Makefile string

func makefile() core.File {
	return core.BuildFile("Makefile", Makefile)
}

//go:embed README.md
var Readme string

func readme() core.File {
	return core.BuildFile("README.md", Readme)
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
