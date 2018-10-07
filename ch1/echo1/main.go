package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(printArgs(os.Args))
}

func printArgs(args []string) string {
	s, sep := "", ""
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func TurboArgsPrint(args []string) string {
	return strings.Join(os.Args[1:], " ")
}
