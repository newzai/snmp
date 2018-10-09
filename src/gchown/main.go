package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("chown", "uftp:uftp", os.Args[1])
	err := cmd.Run()
	fmt.Println(err)
}
