package main

import (
	"fmt"

	. "github.com/colinperry07/DontPassGo/internal/builder"
)

func main() {

	fmt.Println(Build("DontPassGo Password Manager", WithForeground(31)))
}
