package main

import (
	"fmt"

	"github.com/elnaterator/bootstrapper/pkg/core"
	"github.com/elnaterator/bootstrapper/pkg/golang"
	"github.com/elnaterator/bootstrapper/pkg/input"
	"github.com/elnaterator/bootstrapper/pkg/python"
)

func main() {

	fmt.Print("\nBOOTSTRAPPER\n\n")

	project := &core.Project{}
	project.ProjectType = input.Choice("Project type?", []string{"golang", "python"})
	project.RootDir = input.Text("Project root directory?")

	switch project.ProjectType {
	case "golang":
		golang.NewBootstrapper(project).Bootstrap()
	case "python":
		python.NewBootstrapper(project).Bootstrap()
	}

}
