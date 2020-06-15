package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ding <cmd>")
		return
	}

	var cmd *exec.Cmd
	if len(os.Args) > 1 {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	} else {
		cmd = exec.Command(os.Args[1])
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
