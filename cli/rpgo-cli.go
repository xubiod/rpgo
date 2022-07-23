package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"xubiod/rpgo"
)

func main() {
	var extractAllPath string
	var dumpInto string
	var overrideFiles bool

	flag.StringVar(&extractAllPath, "extract", "", "project to extract all files from")
	flag.StringVar(&dumpInto, "o", "", "output directory (can overwrite files)")
	flag.BoolVar(&overrideFiles, "overwrite-files", false, "overwrite existing files")

	flag.Parse()

	if len(extractAllPath) == 0 || len(dumpInto) == 0 {
		fmt.Println("usage: rpgo-cli.go -extract")
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stat(extractAllPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(fmt.Errorf("archive does not exist: %s", extractAllPath).Error())
		os.Exit(1)
	}

	goatVersion := rpgo.GetRPGMakerVersion(extractAllPath)

	switch goatVersion {
	case rpgo.RPGMakerXp, rpgo.RPGMakerVx:
		var goat *rpgo.RGSSADv1

		goat, err = rpgo.MakeRGSSADv1(extractAllPath)

		if err != nil {
			fmt.Printf("error making rgssadv1: %s", err)
			os.Exit(1)
		}

		err = goat.ExtractAllFiles(dumpInto, overrideFiles)

		if err != nil {
			fmt.Printf("error extracting files: %s", err)
			os.Exit(1)
		}

		fmt.Printf("extract completed located at %s", dumpInto)

	case rpgo.RPGMakerVxAce:
		var goat *rpgo.RGSSADv3

		goat, err = rpgo.MakeRGSSADv3(extractAllPath)

		if err != nil {
			fmt.Printf("error making rgssadv3: %s", err)
			os.Exit(1)
		}

		err = goat.ExtractAllFiles(dumpInto, overrideFiles)

		if err != nil {
			fmt.Printf("error extracting files: %s", err)
			os.Exit(1)
		}

		fmt.Printf("extract completed located at %s", dumpInto)

	default:
		fmt.Println("invalid archive")
		os.Exit(1)
	}
}
