package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/v4lomyr/POSIX-shell/internal/parser"
)

var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		cmd = strings.TrimSpace(cmd)

		switch {
		case parser.Exit(cmd):
		case parser.Echo(cmd):
		default:
			fmt.Fprintln(os.Stdout, cmd+": command not found")
		}
	}
}
