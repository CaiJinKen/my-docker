package main

import (
	"os"
	"os/exec"
	"syscall"
	"fmt"
	"log"
)

//docker run bash
func main() {
	str := os.Args[1]
	switch str {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Unknown command.")
	}

}

func run() {
	fmt.Println(os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUTS,
	}
	fmt.Println("F pid is:", syscall.Getpid())

	HandlErr(cmd.Run())
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("S pid is:", syscall.Getpid())

	syscall.Chroot("/home/jin/Desktop/rootfs/")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	HandlErr(cmd.Run())

}

func HandlErr(err error) {
	log.Fatal(err)
}
