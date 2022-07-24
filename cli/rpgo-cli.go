package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"xubiod/rpgo"
)

func main() {
	var input string
	var action string
	var output string
	var overrideFiles bool

	flag.StringVar(&action, "action", "", "action to perform, always required, detailed below")
	flag.StringVar(&input, "i", "", "input (dependent on action)")
	flag.StringVar(&output, "o", "", "output (dependent on action)")
	flag.BoolVar(&overrideFiles, "overwrite-files", false, "overwrite existing files")

	flag.Parse()

	if len(action) == 0 {
		doDefaults()
	}

	goatVersion := rpgo.GetRPGMakerVersion(input)

	switch action {
	case "extract", "decrypt", "dump":
		if len(input) == 0 || len(output) == 0 {
			doDefaults()
		}

		_, err := os.Stat(input)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(fmt.Errorf("archive does not exist: %s", input).Error())
			os.Exit(1)
		}

		switch goatVersion {
		case rpgo.RPGMakerXp, rpgo.RPGMakerVx:
			var goat *rpgo.RGSSADv1

			goat, err = rpgo.NewRGSSADv1(input)

			if err != nil {
				fmt.Printf("error making rgssadv1: %s", err)
				os.Exit(1)
			}

			err = goat.ExtractAllFiles(output, overrideFiles)

			if err != nil {
				fmt.Printf("error extracting files: %s", err)
				os.Exit(1)
			}

			fmt.Printf("extract completed located at %s", output)

		case rpgo.RPGMakerVxAce:
			var goat *rpgo.RGSSADv3

			goat, err = rpgo.NewRGSSADv3(input)

			if err != nil {
				fmt.Printf("error making rgssadv3: %s", err)
				os.Exit(1)
			}

			err = goat.ExtractAllFiles(output, overrideFiles)

			if err != nil {
				fmt.Printf("error extracting files: %s", err)
				os.Exit(1)
			}

			fmt.Printf("extract completed located at %s", output)

		default:
			fmt.Println("invalid archive")
			os.Exit(1)
		}

	case "files", "list", "ls", "dir":
		if len(input) == 0 {
			doDefaults()
		}

		_, err := os.Stat(input)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(fmt.Errorf("archive does not exist: %s", input).Error())
			os.Exit(1)
		}

		var goat *rpgo.RGSSAD
		switch goatVersion {
		case rpgo.RPGMakerXp, rpgo.RPGMakerVx:
			tempgoat, err := rpgo.NewRGSSADv1(input)

			if err != nil {
				fmt.Printf("error making rgssadv1: %s", err)
				os.Exit(1)
			}

			goat = (*rpgo.RGSSAD)(tempgoat)

		case rpgo.RPGMakerVxAce:
			tempgoat, err := rpgo.NewRGSSADv3(input)

			if err != nil {
				fmt.Printf("error making rgssadv3: %s", err)
				os.Exit(1)
			}

			goat = (*rpgo.RGSSAD)(tempgoat)
		default:
			fmt.Println("invalid archive")
			os.Exit(1)
		}

		var szStr string
		var szCompress float64
		szCompressFactor := 1000.0
		var i int
		szStrList := []string{"B", "KB", "MB", "GB"}
		totalBytes := 0.0

		if output == "kibi" {
			szCompressFactor = 1024
			szStrList = []string{"B", "KiB", "MiB", "GiB"}
		} else {
			szStr = "bytes"
		}

		for _, archivefile := range goat.ArchivedFiles {
			szCompress = float64(archivefile.Size)
			totalBytes += szCompress
			i = 0

			if output != "bytes" {
				for szCompress > szCompressFactor {
					i++
					szCompress /= szCompressFactor
				}
				szStr = szStrList[i]
			}

			if i == 0 {
				fmt.Printf("%s\t(%3.0f\t%s)\n", archivefile.Name, szCompress, szStr)
			} else {
				fmt.Printf("%s\t(%3.2f\t%s)\n", archivefile.Name, szCompress, szStr)
			}
		}

		szCompress = totalBytes
		i = 0

		if output != "bytes" {
			for szCompress > szCompressFactor {
				i++
				szCompress /= szCompressFactor
			}
			szStr = szStrList[i]
		}

		fmt.Printf("\ntotal number of files: %d\t(%3.2f\t%s)", len(goat.ArchivedFiles), szCompress, szStr)

	case "justproject":
		if len(input) == 0 || len(output) == 0 {
			doDefaults()
		}

		versionTo := rpgo.RPGMakerInvalid

		switch input {
		case "XP", "RPG Maker XP", "xp", "rmxp", "rgssad", rpgo.XpArchiveName, rpgo.XpProjectFileExtension:
			versionTo = rpgo.RPGMakerXp
		case "VX", "RPG Maker VX", "vx", "rmvx", "rgss2a", rpgo.VxArchiveName, rpgo.VxProjectFileExtension:
			versionTo = rpgo.RPGMakerVx
		case "VXAce", "VX Ace", "RPG Maker VXAce", "RPG Maker VX Ace", "vxace", "vxa", "rmvxace", "rmvxa", "rgss3a", rpgo.VxAceArchiveName, rpgo.VxAceProjectFileExtension:
			versionTo = rpgo.RPGMakerVxAce
		default:
			doDefaults()
		}

		err := rpgo.GenerateProject(versionTo, output)

		if err != nil {
			fmt.Printf("error generating project file:")
		}

		fmt.Println("generated project files")
	}
}

func doDefaults() {
	fmt.Println("usage: rpgo-cli.go -action=[action] [-io] [-overwrite-files]")
	flag.PrintDefaults()
	fmt.Println("\nactions - action flag always required:")
	fmt.Println("\textract/decrypt/dump - extract all files in the archive to the output directory\n\t\ti - input project\n\t\to - output directory\n\t\toverwrite-files - overwrite files toggle")
	fmt.Println("\tfiles/list/ls/dir - list all files in the archive, prints to stdout (use pipes to put into files)\n\t\to - output size format (kilo (default/invalid), kibi, bytes)\n\t\t\tkilo - kilo/megabytes, kibi - kibi/mebibytes, bytes - just bytes")
	fmt.Println("\tjustproject - just generates project files in the requested version to output directory\n\t\ti - version\n\t\t\teither \"xp\"/\"rgssad\", \"vx\"/\"rgss2a\", \"vxace\"/\"rgss3a\" (more variants checked)\n\t\to - output directory")
	os.Exit(1)
}
