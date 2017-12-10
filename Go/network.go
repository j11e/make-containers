package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWNET

	// Run waits for the Command to finish. Alternative is Start then Wait
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}