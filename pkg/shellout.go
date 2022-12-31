package shellout

import (
	"os"
	"os/exec"
	"io/ioutil"
	"io"
	"fmt"
)

func Cmd(cmdStart []string, script string, args ...string) (cmd *exec.Cmd, path string, err error) {
	if len(cmdStart) < 1 {
		return nil, "", fmt.Errorf("GenericCmd: cmdStart empty")
	}

	file, err := ioutil.TempFile(".", "shell*.sh")
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	fmt.Fprint(file, script)
	path = file.Name()
	cargs := append([]string{}, cmdStart[1:]...)
	cargs = append(cargs, path)
	cargs = append(cargs, args...)

	cmd = exec.Command(cmdStart[0], cargs...)
	return cmd, path, err
}

func Out(cmdStart []string, script string, args ...string) error {
	cmd, path, err := Cmd(cmdStart, script, args...)
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

func Piped(cmdStart []string, script string, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) error {
	cmd, path, err := Cmd(cmdStart, script, args...)
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

func ShellCmd(script string, args ...string) (cmd *exec.Cmd, path string, err error) {
	return Cmd([]string{"bash"}, script, args...)
}
func ShellOut(script string, args ...string) error {
	return Out([]string{"bash"}, script, args...)
}
func ShellPiped(script string, stdin io.Reader, stdout io.Writer, stderr io.Writer, args ...string) error {
	return Piped([]string{"bash"}, script, stdin, stdout, stderr, args...)
}
