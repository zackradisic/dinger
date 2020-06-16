package main

import (
	"flag"
	"os"

	"github.com/zackradisic/dinger/cmd"
)

var (
	sound = flag.String("sound",
		os.Getenv("GOPATH")+"/src/github.com/zackradisic/dinger/sounds/ding.mp3",
		"The path to the sound file to play")
)

func main() {
	cmd.Execute()
}
