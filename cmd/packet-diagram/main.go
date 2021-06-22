package main

import (
	"log"
	"os"

	packetdiagram "github.com/bitbears-dev/packet-diagram"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	InputFile string `short:"i" long:"input" required:"true"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("error: %#v\n", err)
	}
}

func run() error {
	_, err := flags.Parse(&opts)
	if err != nil {
		return err
	}

	f, err := os.Open(opts.InputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	def, err := packetdiagram.LoadDefinition(f)
	if err != nil {
		return err
	}

	w := os.Stdout
	err = packetdiagram.Draw(def, w)
	if err != nil {
		return err
	}

	return nil
}
