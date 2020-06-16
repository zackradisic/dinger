package cmd

import (
	"fmt"
	"os"
)

// Invoker is the struct that handles invoking commands
type Invoker struct {
	cmds map[string]*command
}

// Run executes the given command
func (i *Invoker) Run() error {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:%s\n", i.usageString())
		return nil
	}

	if c, ok := i.cmds[os.Args[1]]; ok {
		err := c.execute(os.Args[2:])
		if err != nil {
			return err
		}

		return nil
	}

	fmt.Printf("%s\n", i.unknownCommandString(os.Args[0]))
	return nil
}

func (i *Invoker) init() {
	i.registerCommand("run", newRunCommand())
}

func (i *Invoker) registerCommand(name string, cmd *command) {
	i.cmds[name] = cmd
}

func (i *Invoker) unknownCommandString(cmdName string) string {
	return fmt.Sprintf("Error: unknown command \"%s\"", cmdName)
}

func (i *Invoker) usageString() string {
	s := "Usage:"
	for _, cmd := range i.cmds {
		s += "\n  " + usageString(cmd)
	}

	return s
}

// CreateInvoker creates a command invoker
func CreateInvoker() *Invoker {
	invoker := &Invoker{
		cmds: make(map[string]*command),
	}

	invoker.init()
	return invoker
}
