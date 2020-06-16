package cmd

import "fmt"

type command struct {
	execute      func(args []string) error
	args         []string
	validateArgs func(args []string) error
}

func usageString(args []string) string {
	if len(args) == 0 {
		return ""
	} else if len(args) == 1 {
		return fmt.Sprintf("<%s>", args[0])
	} else {
		s := fmt.Sprintf("<%s", args[0])
		for _, arg := range args {
			s += fmt.Sprintf(" | %s", arg)
		}
		s += ">"
		return s
	}
}

func contains(args []string, arg string) bool {
	for _, a := range args {
		if a == arg {
			return true
		}
	}

	return false
}
