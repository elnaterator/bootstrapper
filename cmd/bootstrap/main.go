package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/elnaterator/bootstrapper/pkg/color"
	"github.com/elnaterator/bootstrapper/pkg/core"
	"github.com/elnaterator/bootstrapper/pkg/golang"
	"github.com/elnaterator/bootstrapper/pkg/input"
	"github.com/elnaterator/bootstrapper/pkg/python"
)

func main() {

	fmt.Print("\nBOOTSTRAPPER\n\n")

	project := &core.Project{}

	rootDir := input.Text("Project root directory?")
	project.SetRootDir(rootDir)

	projectType := input.Choice("Project type?", ProjectTypes)
	project.SetType(projectType)

	language := input.Choice("Language?", LanguagesByProjectType[projectType])
	project.SetLanguage(language)

	buildSystem := input.Choice("Build system?", BuildSystemsByLanguage[language])
	project.SetBuildSystem(buildSystem)

	b := newBootstrapper(project)
	b.CollectAdditionalOptions()

	dir := b.BuildModel()
	fmt.Println(color.Term("Project files to be created:", color.Primary))
	renderDir(dir, 0)

	confirm := input.Choice("Confirm project creation?", []string{"yes", "no"})
	if confirm != "yes" {
		fmt.Println(color.Term("Skipping project creation.", color.Primary))
	} else {
		fmt.Println(color.Term("Creating project...", color.Primary))
		b.Bootstrap()
		fmt.Println(color.Term("Project created.", color.Primary))
	}

}

var ProjectTypes = []string{"basic", "cli"}

var LanguagesByProjectType = map[string][]string{"basic": {"golang", "python"}, "cli": {"golang", "python"}}

var BuildSystemsByLanguage = map[string][]string{"golang": {"gomod"}, "python": {"poetry", "pip"}}

func newBootstrapper(p *core.Project) core.Bootstrapper {
	switch p.Language {
	case "golang":
		return golang.NewBootstrapper(p)
	case "python":
		return python.NewBootstrapper(p)
	}
	log.Fatal("no bootstrapper for given project options")
	return nil
}

func renderDir(dir *core.Directory, depth int) {

	if dir == nil {
		log.Fatal("invalid directory structure")
	}

	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s%s/\n", indent, dir.Name)

	for _, f := range dir.Files {
		fmt.Printf("%s  %s\n", indent, f.Name)
	}

	for _, d := range dir.Directories {
		renderDir(&d, depth+1)
	}

}
