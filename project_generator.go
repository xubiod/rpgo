package rpgo

// RPGMakerDecrypter.Decrypter/ProjectGenerator.cs

import (
	"fmt"
	"os"
	"path"
)

func GenerateProject(version RPGMakerVersion, outputPath string) {
	var content string
	var extension string
	var ini string

	switch version {
	case RPGMakerXp:
		content = XpProjectFileContent
		extension = XpProjectFileExtension
		ini = XpIniFileContents
	case RPGMakerVx:
		content = VxProjectFileContent
		extension = VxProjectFileExtension
		ini = VxIniFileContents
	case RPGMakerVxAce:
		content = VxAceProjectFileContent
		extension = VxAceProjectFileExtension
		ini = VxAceIniFileContents
	default:
		content = ""
		extension = ""
		ini = ""
	}

	file, err := os.Create(path.Join(outputPath, fmt.Sprintf("Game.%s", extension)))

	if err != nil {
		panic("a")
	}

	_, err = file.WriteString(content)

	if err != nil {
		file.Close()
		panic("b")
	}

	file.Close()

	file, err = os.Create(path.Join(outputPath, "Game.ini"))

	if err != nil {
		panic("c")
	}

	_, err = file.WriteString(ini)

	if err != nil {
		file.Close()
		panic("d")
	}

	file.Close()
}
