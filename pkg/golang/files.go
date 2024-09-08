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

//go:embed gitignore
var Gitignore string

func gitignore() core.File {
	return core.BuildFile(".gitignore", Gitignore)
}
