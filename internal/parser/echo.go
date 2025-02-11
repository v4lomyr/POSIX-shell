package parser

import (
	"fmt"
	"os"
	"regexp"
)

func Echo(cmd string) bool {
	re := regexp.MustCompile(`^echo(?:\s+(.*))?$`)
	if !re.MatchString(cmd) {
		return false
	}

	matches := re.FindStringSubmatch(cmd)
	text := ""
	if len(matches) > 0 {
		text = matches[len(matches)-1]
		fmt.Fprintln(os.Stdout, text)
	}

	return true
}
