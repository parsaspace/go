package main

import (
	"fmt"

	parsaspace "github.com/parsaspace/go"
)

func main() {
	p := parsaspace.NewClient("TOKEN")

	err := p.Upload("yourdomain.parsaspace.com", "/", "main.go")

	if err != nil {
		fmt.Printf("%s", err)
	}

	ls, err := p.Files("yourdomain.parsaspace.com", "/")
	if err != nil {
		fmt.Printf("%s", err)
	}

	for _, v := range ls {
		fmt.Println(v.Name)
	}
}
