package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 执行一个任务

	cmd := exec.Command("/bin/bash", "-c", "sleep 5; ls -l")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
