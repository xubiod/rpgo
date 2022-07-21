package rpgo

// RPGMakerDecrypter.Decryptre/RGSSAD.cs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

type RGSSAD struct {
	Filepath      string
	ByteReader    bytes.Reader
	ArchivedFiles []RPGMakerArchivedFile
}

func (RGSSAD) Make(filepath string) *RGSSAD {
	created := new(RGSSAD)
	created.Filepath = filepath
	return created
}

func (rpg *RGSSAD) GetVersion() (RPGMakerVersion, error) {
	var header string

	header, err := ReadCString(rpg.ByteReader, 7)

	if header != RGSSADHeader || err != nil {
		return RPGMakerInvalid, err
	}

	result, err := rpg.ByteReader.ReadByte()

	return RPGMakerVersion(result), err
}

// Currently overwrite existing is ignored
func (rpg *RGSSAD) ExtractFile(archivedFile RPGMakerArchivedFile, outputDirectoryPath string, overwriteExisting bool, createDirectory bool) error {
	var outputPath string

	if createDirectory {
		directoryPath := path.Dir(archivedFile.Name)

		if directoryPath == "." {
			return errors.New("rpgo/rgssad: invalid file path")
		}

		outputPath = path.Join(outputDirectoryPath, directoryPath)
	} else {
		splitted := strings.Split(archivedFile.Name, "\\")
		filename := splitted[len(splitted)-1]

		outputPath = path.Join(outputDirectoryPath, filename)
	}

	rpg.ByteReader.Seek(archivedFile.Offset, io.SeekStart)
	data := make([]byte, archivedFile.Size)
	rpg.ByteReader.Read(data)

	finalFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	finalFile.Write(rpg.decryptFileData(data, uint32(archivedFile.Key)))

	return finalFile.Close()
}

// Currently overwrite existing is ignored
func (rpg *RGSSAD) ExtractAllFiles(outputDirectoryPath string, overrideExisting bool) {
	for _, archivedFile := range rpg.ArchivedFiles {
		rpg.ExtractFile(archivedFile, outputDirectoryPath, overrideExisting, true)
	}
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
