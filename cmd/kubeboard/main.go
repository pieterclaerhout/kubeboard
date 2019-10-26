package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/kubeboard"
)

func main() {

	file, _ := os.Create("/Users/pclaerhout/Desktop/kubectl.log")
	defer file.Close()

	log.PrintTimestamp = true
	log.Stdout = file
	log.Stderr = file

	kb := kubeboard.NewKubeBoard()

	go func() {

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-signalChan

		kb.Stop()

	}()

	kb.Start()

}
