package main

import (
	"os"
	"os/exec"
	"io/ioutil"
	"io"
	"fmt"

	"strings"
)

func ShellCmd(script string) (cmd *exec.Cmd, path string, err error) {
	file, err := ioutil.TempFile(".", ".shell")
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	fmt.Fprint(file, script)
	path = file.Name()

	cmd = exec.Command("bash", path)
	return cmd, path, err
}

func ShellOut(script string) error {
	cmd, path, err := ShellCmd(script)
	defer os.Remove(path)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ShellOutPiped(script string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	cmd, path, err := ShellCmd(script)
	defer os.Remove(path)
	if err != nil {
		return err
	}
	if stdin != nil {
		cmd.Stdin = stdin
	}
	if stdout != nil {
		cmd.Stdout = stdout
	}
	if stderr != nil {
		cmd.Stderr = stderr
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var b strings.Builder
	err := ShellOutPiped("ls -ltr", nil, &b, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())

	b.Reset()
	r := strings.NewReader("5\n10\n15\n20")
	err = ShellOutPiped(`awk '{print $1+3}'`, r, &b, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())
}
