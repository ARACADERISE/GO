package main

import (
	"fmt"
	ps"package/src"
)

func main() {
	info := ps.Gather_default_file("main.go")

	fmt.Println(info)
}
