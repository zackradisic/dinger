package cmd

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/zackradisic/dinger/config"
)

func newRunCommand() *command {
	run := &command{
		name:         "run",
		execute:      executeRunCommand,
		args:         []string{"cmd"},
		validateArgs: validateRunCommandArgs,
	}

	return run
}

func executeRunCommand(args []string) error {
	var command *exec.Cmd
	if len(args) > 1 {
		command = exec.Command(args[0], args[1:]...)
	} else {
		command = exec.Command(args[0])
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	sound := config.Configuration.Sound
	f, err := os.Open(sound)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}

func validateRunCommandArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("No command specified")
	}

	return nil
}
