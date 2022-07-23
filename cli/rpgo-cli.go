package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"xubiod/rpgo"
)

func main() {
	var inputPath string
	var action string
	var dumpInto string
	var overrideFiles bool

	flag.StringVar(&inputPath, "i", "", "project to do action")
	flag.StringVar(&action, "action", "", "action (extract/decrypt/dump, files/list/ls)")
	flag.StringVar(&dumpInto, "o", "", "output directory")
	flag.BoolVar(&overrideFiles, "overwrite-files", false, "overwrite existing files")

	flag.Parse()

	if len(action) == 0 {
		doDefaults()
	}

	_, err := os.Stat(inputPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(fmt.Errorf("archive does not exist: %s", inputPath).Error())
		os.Exit(1)
	}

	goatVersion := rpgo.GetRPGMakerVersion(inputPath)

	switch action {
	case "extract", "decrypt", "dump":
		if len(inputPath) == 0 || len(dumpInto) == 0 {
			doDefaults()
		}

		switch goatVersion {
		case rpgo.RPGMakerXp, rpgo.RPGMakerVx:
			var goat *rpgo.RGSSADv1

			goat, err = rpgo.MakeRGSSADv1(inputPath)

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

			goat, err = rpgo.MakeRGSSADv3(inputPath)

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

	case "files", "list", "ls":
		if len(inputPath) == 0 {
			doDefaults()
		}
		switch goatVersion {
		case rpgo.RPGMakerXp, rpgo.RPGMakerVx:
			var goat *rpgo.RGSSADv1

			goat, err = rpgo.MakeRGSSADv1(inputPath)

			if err != nil {
				fmt.Printf("error making rgssadv1: %s", err)
				os.Exit(1)
			}

			for _, archivefile := range goat.ArchivedFiles {
				fmt.Printf("%s\t(%d bytes)\n", archivefile.Name, archivefile.Size)
			}

		case rpgo.RPGMakerVxAce:
			var goat *rpgo.RGSSADv3

			goat, err = rpgo.MakeRGSSADv3(inputPath)

			if err != nil {
				fmt.Printf("error making rgssadv3: %s", err)
				os.Exit(1)
			}

			for _, archivefile := range goat.ArchivedFiles {
				fmt.Printf("%s\t(%d bytes)\n", archivefile.Name, archivefile.Size)
			}

		default:
			fmt.Println("invalid archive")
			os.Exit(1)
		}
	}
}

func doDefaults() {
	fmt.Println("usage: rpgo-cli.go -i=[archive file] -action=[action] -o=[directory]")
	flag.PrintDefaults()
	os.Exit(1)
}
