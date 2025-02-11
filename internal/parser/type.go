package parser

import (
	"fmt"
	"github.com/v4lomyr/POSIX-shell/internal/constant"
	"os"
	"regexp"
	"slices"
)

func Type(cmd string) bool {
	re := regexp.MustCompile(`^type\s+(.*)$`)

	if !re.MatchString(cmd) {
		return false
	}

	matches := re.FindStringSubmatch(cmd)
	payload := matches[len(matches)-1]

	if slices.Contains(constant.CommandList, payload) {
		fmt.Fprintln(os.Stdout, payload+" is a shell builtin")
	} else {
		fmt.Fprintln(os.Stdout, payload+": not found")
	}

	return true
}
