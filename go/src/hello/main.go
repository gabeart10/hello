package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
)

func plain(text string) {
	fmt.Println(c.Clear + c.Rc() + text + c.Reset)
}

func color(message string) {
	fmt.Print(c.Clear)
	for {
		fmt.Print(c.Rc() + message + c.Reset)
	}
}
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "color" {
			color("Hello World")
		} else {
			plain("Hello World")
		}
	} else {
		plain("Hello World")
	}
}
