package main

import (
	"fmt"
	"log"

	"github.com/elnaterator/bootstrapper/pkg/color"
	"github.com/elnaterator/bootstrapper/pkg/golang"
	"github.com/elnaterator/bootstrapper/pkg/options"
	"github.com/elnaterator/bootstrapper/pkg/python"
)

func main() {

	fmt.Print("\nBOOTSTRAPPER\n\n")

	var dir string
	fmt.Println(color.Wrap("project directory?", color.Red))
	_, err := fmt.Scan(&dir)
	if err != nil {
		log.Fatal("unable to receive input")
	}
	fmt.Println(dir)

	lang := options.GetOptions("project type? ", []string{"golang", "python"})

	switch lang {
	case "golang":
		golang.NewBootstrapper().Bootstrap()
	case "python":
		python.NewBootstrapper().Bootstrap()
	}

}
