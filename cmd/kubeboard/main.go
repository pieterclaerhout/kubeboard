package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pieterclaerhout/kubeboard"
)

func main() {

	kb := kubeboard.NewKubeBoard()

	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		s := <-c

		// The signal is received, you can now do the cleanup
		fmt.Println("Got signal:", s)
		kb.Stop()

	}()

	kb.Start()

}
