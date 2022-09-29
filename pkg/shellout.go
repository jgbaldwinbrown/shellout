package shellout

import (
	"os"
	"os/exec"
	"io/ioutil"
	"io"
	"fmt"
)

func ShellCmd(script string, args ...string) (cmd *exec.Cmd, path string, err error) {
	file, err := ioutil.TempFile(".", ".shell")
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	fmt.Fprint(file, script)
	path = file.Name()
	cargs := append([]string{path}, args...)

	cmd = exec.Command("bash", cargs...)
	return cmd, path, err
}

func ShellOut(script string, args ...string) error {
	cmd, path, err := ShellCmd(script, args...)
	defer os.Remove(path)
	if err != nil {
		return err
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ShellOutPiped(script string, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) error {
	cmd, path, err := ShellCmd(script, args...)
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
