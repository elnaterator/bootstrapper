package main

import (
	"fmt"
	"log"

	"github.com/elnaterator/bootstrapper/pkg/greet"
	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{}
	greeter := &greet.Greeter{Name: "nate"}
	greeting := greeter.BuildGreeting()
	fmt.Println(greeting)

	err := app.Run([]string{})
	if err != nil {
		log.Fatalf("Error running bootstrapper: %v", err)
	}

}
