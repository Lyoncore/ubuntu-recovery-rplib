package rplib

import (
	"os"
	"os/exec"
	"strings"

	"github.com/snapcore/snapd/logger"
)

func Shellexec(name string, args ...string) {
	logger.Debugf(name, args)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	Checkerr(err)
}

func Shellexecoutput(name string, args ...string) string {
	logger.Debugf(name, args)
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	Checkerr(err)

	return strings.TrimSpace(string(out[:]))
}

func Shellcmd(command string) {
	cmd := exec.Command("sh", "-c", command)
	logger.Debugf(strings.Join(cmd.Args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	Checkerr(err)
}

func Shellcmdoutput(command string) string {
	cmd := exec.Command("sh", "-c", command)
	logger.Debugf(strings.Join(cmd.Args, " "))
	out, err := cmd.Output()
	Checkerr(err)

	return strings.TrimSpace(string(out[:]))
}
