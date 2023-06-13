package rpgo

// RPGMakerDecrypter.Decryptre/RGSSAD.cs

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// RGSSAD
//
// The generic structure of a RPG Maker encrypted archive.
type RGSSAD struct {
	Filepath      string
	Data          []byte
	ByteReader    bytes.Reader
	ArchivedFiles []RPGMakerArchivedFile
}

// NewRGSSAD
//
// Creates a new RGSSAD structure and configures it for use.
//
// Returns a pointer to the created structure.
func NewRGSSAD(filepath string) *RGSSAD {
	created := new(RGSSAD)
	created.Filepath = filepath

	f, _ := os.Open(created.Filepath)

	var err error
	created.Data, err = io.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	created.ByteReader = *bytes.NewReader(created.Data)

	return created
}

// GetVersion
//
// Gets the RPGMakerVersion of the encrypted archive.
//
// Returns the version and nil for error on success; RPGMakerInvalid and an
// error otherwise.
func (rpg *RGSSAD) GetVersion() (RPGMakerVersion, error) {
	var header string

	header, err := ReadCString(&rpg.ByteReader, 7)

	if header != RGSSADHeader || err != nil {
		return RPGMakerInvalid, err
	}

	result, err := rpg.ByteReader.ReadByte()

	return RPGMakerVersion(result), err
}

// ExtractFile
//
// Extracts the given archived file from the encrypted archive.
//
// Returns nil on success, error otherwise.
func (rpg *RGSSAD) ExtractFile(archivedFile RPGMakerArchivedFile, outputDirectoryPath string, overwriteExisting bool, createDirectory bool) error {
	var outputPath string

	if createDirectory {
		subDirectories := strings.Split(archivedFile.Name, string(filepath.Separator))
		subDirectories = subDirectories[:len(subDirectories)-1]

		outputPath = outputDirectoryPath

		for _, itm := range subDirectories {
			outputPath = filepath.Join(outputPath, itm)
			err := os.Mkdir(outputPath, os.ModeDir)
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

// ExtractAllFiles
//
// Extracts all files from the archive.
//
// Returns nil on success, error otherwise.
func (rpg *RGSSAD) ExtractAllFiles(outputDirectoryPath string, overrideExisting bool) (err error) {
	for _, archivedFile := range rpg.ArchivedFiles {
		err := rpg.ExtractFile(archivedFile, outputDirectoryPath, overrideExisting, true)
		if err != nil {
			return err
		}
	}

	return nil
}

// Decrypts the file data from a byte slice into another byte slice.
//
// This function is meant for internal use only in ExtractFile.
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
