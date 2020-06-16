package cmd

import "fmt"

type command struct {
	execute      func(args []string) error
	args         []string
	validateArgs func(args []string) error
}

func usageString(c *command) string {
	if len(c.args) == 0 {
		return ""
	} else if len(c.args) == 1 {
		return fmt.Sprintf("<%s>", c.args[0])
	} else {
		s := fmt.Sprintf("<%s", c.args[0])
		for _, arg := range c.args {
			s += fmt.Sprintf(" | %s", arg)
		}
		s += ">"
		return s
	}
}
