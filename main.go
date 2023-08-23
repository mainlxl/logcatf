package main

import (
	"os"

	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
	log.SetOutput(colorable.NewColorableStderr())
}

func main() {
	cli := &CLI{
		inStream:  os.Stdin,
		outStream: colorable.NewColorableStdout(),
		errStream: colorable.NewColorableStderr(),
	}
	os.Exit(cli.Run(os.Args))
}
