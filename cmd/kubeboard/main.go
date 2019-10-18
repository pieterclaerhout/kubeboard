package main

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/pieterclaerhout/kubeboard"
)

func main() {
	kb := kubeboard.NewKubeBoard()
	err := kb.Start()
	log.CheckError(err)
}
