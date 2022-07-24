package rpgo

// RPGMakerDecrypter.Decrypter/ProjectGenerator.cs

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// Generates project files appropriate for the given version.
//
// Returns nil on success, error otherwise.
func GenerateProject(version RPGMakerVersion, outputPath string) error {
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
		return errors.New("rpgo/generateproject: no valid archive version")
	}

	file, err := os.Create(path.Join(outputPath, fmt.Sprintf("Game.%s", extension)))

	if err != nil {
		return err
	}

	_, err = file.WriteString(content)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()

	file, err = os.Create(path.Join(outputPath, "Game.ini"))

	if err != nil {
		return err
	}

	_, err = file.WriteString(ini)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()

	return nil
}
