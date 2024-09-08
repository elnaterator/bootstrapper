package input

import (
	"fmt"
	"log"

	"github.com/elnaterator/bootstrapper/pkg/color"
)

func Text(title string) string {
	var dir string
	fmt.Println(color.Term(title, color.Primary))
	_, err := fmt.Scan(&dir)
	if err != nil {
		log.Fatal("unable to receive input")
	}
	return dir
}
