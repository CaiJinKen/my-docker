package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
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

	fmt.Println("S pid is:", syscall.Getpid())

	syscall.Chroot("/home/jin/Desktop/rootfs/")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	path, err := exec.LookPath(os.Args[2])
	HandlErr(err)
	HandlErr(syscall.Exec(path, append([]string{path}, os.Args[3:]...), os.Environ()))
}

func HandlErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
