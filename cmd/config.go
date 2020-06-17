package cmd

import (
	"fmt"

	"github.com/zackradisic/dinger/config"
)

func newConfigCommand() *command {
	return &command{
		name:         "config",
		execute:      executeConfigCommand,
		validateArgs: validateRunCommandArgs,
		args:         []string{"set", "get"},
	}
}

func executeConfigCommand(args []string) error {

	if len(args) == 3 {
		switch args[0] {
		case "set":
			err := config.SetValue(args[1], args[2])
			if err != nil {
				return err
			}
		}

		fmt.Printf("Success: %s has been set to %s\n", args[1], args[2])
		return nil
	}

	if len(args) == 2 {
		switch args[0] {
		case "get":
			err := config.PrintValue(args[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateConfigCommandArgs(args []string) error {
	if len(args) == 3 {
		if !contains(args, args[0]) {
			return fmt.Errorf("Unknown argument: \"%s\"", args[0])
		}

		return nil
	}

	return fmt.Errorf("%s", usageString(args))
}
