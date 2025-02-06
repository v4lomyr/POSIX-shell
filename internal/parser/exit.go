package parser

import (
	"os"
	"regexp"
	"strconv"
)

func Exit(cmd string) bool {
	re := regexp.MustCompile(`^exit(?: (\d+))?$`)
	if re.MatchString(cmd) {
		return false
	}

	matches := re.FindStringSubmatch(cmd)
	exitCode := 0
	if len(matches) > 0 {
		exitCode, _ = strconv.Atoi(matches[1])
	}

	os.Exit(exitCode)

	return true
}
