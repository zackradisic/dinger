package main

import (
	"flag"
	"os"

	"github.com/zackradisic/dinger/cmd"
	"github.com/zackradisic/dinger/config"
)

var (
	sound = flag.String("sound",
		os.Getenv("GOPATH")+"/src/github.com/zackradisic/dinger/sounds/ding.mp3",
		"The path to the sound file to play")
)

func main() {
	config.ReadConfig()
	cmdInvoker := cmd.CreateInvoker()

	cmdInvoker.Run()
}
