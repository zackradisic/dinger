package cmd

import (
	"fmt"
	"os"
)

// Invoker is the struct that handles invoking commands
type Invoker struct {
	cmds map[string]*command
}

func (i *Invoker) run() error {
	if len(os.Args) < 1 {
		fmt.Printf("Usage:%s\n", i.usageString())
		return nil
	}

	if c, ok := i.cmds[os.Args[0]]; ok {
		err := c.execute(os.Args[1:])
		if err != nil {
			return err
		}
	}

	fmt.Printf("%s\n", i.unknownCommandString(os.Args[0]))
	return nil
}

func (i *Invoker) init() {

}

func (i *Invoker) registerCommand(name string, cmd *command) {
	i.cmds[name] = cmd
}

func (i *Invoker) unknownCommandString(cmdName string) string {
	return fmt.Sprintf("Error: unknown command \"%s\"", cmdName)
}

func (i *Invoker) usageString() string {
	s := "Usage:"
	for cmdName, cmd := range i.cmds {
		s += "\n  " + usageString(cmd)
	}

	return s
}
