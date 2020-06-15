package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ding <cmd>")
		return
	}

	var cmd *exec.Cmd
	if len(os.Args) > 1 {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	} else {
		cmd = exec.Command(os.Args[1])
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Open("sounds/ding.mp3")
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
}
