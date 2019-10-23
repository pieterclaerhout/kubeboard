package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pieterclaerhout/kubeboard"
)

func main() {

	file, _ := os.Create("/Users/pclaerhout/Desktop/kubectl.log")
	defer file.Close()

	kb := kubeboard.NewKubeBoard()
	kb.File = file

	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		s := <-c

		// The signal is received, you can now do the cleanup
		kb.File.WriteString("Got signal: " + fmt.Sprintf("%v", s) + "\n")
		kb.Stop()

	}()

	kb.Start()

}
