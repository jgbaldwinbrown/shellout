package main

import (
	"github.com/jgbaldwinbrown/shellout/pkg"
	"os"
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	err := shellout.ShellOutPiped("ls -ltr", os.Stdin, &b, os.Stderr)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())

	b.Reset()
	r := strings.NewReader("5\n10\n15\n20")
	err = shellout.ShellOutPiped(`awk '{print $1+3}'`, r, &b, os.Stderr)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())

	err = shellout.ShellOut(`ls -ltrh`)
	if err != nil {
		panic(err)
	}

	err = shellout.ShellOut(`ls "$@"`, "-l", "-t", "-r", "-h")
	if err != nil {
		panic(err)
	}
}
