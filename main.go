package main

import (
	"fmt"

	. "github.com/colinperry07/DontPassGo/internal/builder"
)

func main() {
	fmt.Println(Build("Hello World!", WithForeground(31)))
}
