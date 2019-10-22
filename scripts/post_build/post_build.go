package main

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	args, err := james.ParseBuildArgs()
	log.CheckError(err)

	if args.GOOS == "darwin" {

		macAppPackage := james.NewMacAppPackage(
			args.OutputPath,
			filepath.Join(args.ProjectPath, "icon.png"),
		)
		err := macAppPackage.Create()
		log.CheckError(err)

	}

}
