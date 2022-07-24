package rpgo

// RPGMakerDecrypter.Decryptre/RGSSAD.cs

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type RGSSAD struct {
	Filepath      string
	Data          []byte
	ByteReader    bytes.Reader
	ArchivedFiles []RPGMakerArchivedFile
}

func MakeRGSSAD(filepath string) *RGSSAD {
	created := new(RGSSAD)
	created.Filepath = filepath

	f, _ := os.Open(created.Filepath)

	var err error
	created.Data, err = ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	created.ByteReader = *bytes.NewReader(created.Data)

	return created
}

func (rpg *RGSSAD) GetVersion() (RPGMakerVersion, error) {
	var header string

	header, err := ReadCString(&rpg.ByteReader, 7)

	if header != RGSSADHeader || err != nil {
		return RPGMakerInvalid, err
	}

	result, err := rpg.ByteReader.ReadByte()

	return RPGMakerVersion(result), err
}

func (rpg *RGSSAD) ExtractFile(archivedFile RPGMakerArchivedFile, outputDirectoryPath string, overwriteExisting bool, createDirectory bool) error {
	var outputPath string

	if createDirectory {
		// directoryPath := path.Dir(strings.Replace(archivedFile.Name, "\\", "/", -1))

		// if directoryPath == "." {
		// 	return errors.New("rpgo/rgssad: invalid file path")
		// }

		subDirectories := strings.Split(archivedFile.Name, string(filepath.Separator))
		subDirectories = subDirectories[:len(subDirectories)-1]

		outputPath = outputDirectoryPath

		for _, itm := range subDirectories {
			outputPath = filepath.Join(outputPath, itm)
			err := os.Mkdir( /*filepath.Dir(*/ outputPath, os.ModeDir)
			_, err2 := os.Stat(filepath.Dir(outputPath))

			if os.IsNotExist(err) && os.IsNotExist(err2) {
				return err
			}
		}

		outputPath = filepath.Join(outputPath, strings.Split(archivedFile.Name, string(filepath.Separator))[len(subDirectories)])
	} else {
		splitted := strings.Split(archivedFile.Name, "\\")
		filename := splitted[len(splitted)-1]

		outputPath = filepath.Join(outputDirectoryPath, filename)
	}

	rpg.ByteReader.Seek(archivedFile.Offset, io.SeekStart)
	data := make([]byte, archivedFile.Size)
	rpg.ByteReader.Read(data)

	if _, err := os.Stat(outputPath); os.IsNotExist(err) || overwriteExisting {
		finalFile, err := os.Create(outputPath)

		if err != nil {
			return err
		}

		finalFile.Write(rpg.decryptFileData(data, uint32(archivedFile.Key)))

		return finalFile.Close()
	}

	return nil
}

func (rpg *RGSSAD) ExtractAllFiles(outputDirectoryPath string, overrideExisting bool) error {
	for _, archivedFile := range rpg.ArchivedFiles {
		err := rpg.ExtractFile(archivedFile, outputDirectoryPath, overrideExisting, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (*RGSSAD) decryptFileData(encryptedFileData []byte, key uint32) []byte {
	decryptedFileData := make([]byte, len(encryptedFileData))

	tempKey := key
	keyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(keyBytes, key)
	i := 0
	j := 0

	for i < len(encryptedFileData) {
		if j == 4 {
			j = 0
			tempKey *= 7
			tempKey += 3
			binary.LittleEndian.PutUint32(keyBytes, tempKey)
		}

		decryptedFileData[i] = byte(encryptedFileData[i] ^ keyBytes[j])

		i++
		j++
	}

	return decryptedFileData
}
