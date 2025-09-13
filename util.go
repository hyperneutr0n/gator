package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func printLineLimitter() {
	length := 30
	shellCmd := exec.Command("stty", "size")
	shellCmd.Stdin = os.Stdin
	out, err := shellCmd.Output()
	if err != nil {
		fmt.Println("failed getting terminal width")
	}

	terminalSize := strings.Split(strings.TrimSpace(string(out)), " ")
	width, err := strconv.Atoi(terminalSize[1])
	if err != nil {
		fmt.Println("failed parsing width")
	}

	if width != 0 {
		length = width
	}
	line := ""
	for range length {
		line += "="
	}
	fmt.Println(line)
}