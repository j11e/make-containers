package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	command := strings.Join(os.Args[1:], " ")
	command = "mount --make-rprivate / ; mount -t tmpfs tmpfs /mytmp ; " + command

	fmt.Printf("Executing %s\n", command)
	//return

	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWNS

	// Run waits for the Command to finish. Alternative is Start then Wait
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
