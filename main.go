package main

import (
	"fmt"
)

func main() {
	fmt.Println(modifyText("Hello World!", "31"))
}

func modifyText (textToModify string, modifier string) string {
	prefix := "\033["
	reset := "\033[0m"

	return prefix + modifier + "m" + textToModify + reset
}
