package rpgo

import (
	"encoding/binary"
	"fmt"
	"io"
)

//RPGMakerDecrypter.Decrypter/RGSSADv3.cs

type RGSSADv3 RGSSAD

// NewRGSSADv3
//
// Creates a new RGSSADv3 structure and configures it for use.
//
// Returns a pointer to the created structure, nil on success, nil and error
// otherwise.
func NewRGSSADv3(filepath string) (created *RGSSADv3, err error) {
	created = (*RGSSADv3)(NewRGSSAD(filepath))

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv3 || err != nil {
		return nil, fmt.Errorf("rpgo/rgssadv3: version not v3 or this:\n%s", err.Error())
	}

	created.readRGSSAD()

	return created, nil
}

// Reads the encrypted RGSSADv3 archive and generates a ArchivedFile slice for
// the RGSSAD structure.
//
// This function is meant for internal use in NewRGSSADv3.
func (rpg *RGSSADv3) readRGSSAD() {
	rpg.ByteReader.Seek(8, io.SeekStart)

	t := make([]byte, 4)
	rpg.ByteReader.Read(t)
	num := int(binary.LittleEndian.Uint32(t))

	key := uint(num)

	key *= 9
	key += 3

	for {
		newArchivedFile := new(RPGMakerArchivedFile)

		// OFFSET
		rpg.ByteReader.Read(t)
		num = int(binary.LittleEndian.Uint32(t))

		newArchivedFile.Offset = int64(rpg.decryptInteger(num, key))

		// SIZE
		rpg.ByteReader.Read(t)
		num = int(binary.LittleEndian.Uint32(t))

		newArchivedFile.Size = rpg.decryptInteger(num, key)

		// KEY
		rpg.ByteReader.Read(t)
		num = int(binary.LittleEndian.Uint32(t))

		newArchivedFile.Key = uint(rpg.decryptInteger(num, key))

		if newArchivedFile.Offset == 0 {
			break
		}

		// NAME
		rpg.ByteReader.Read(t)
		num := int(binary.LittleEndian.Uint32(t))

		nameLen := rpg.decryptInteger(num, key)

		u := make([]byte, nameLen)

		rpg.ByteReader.Read(u)
		newArchivedFile.Name = rpg.decryptFilename(u, key)

		rpg.ArchivedFiles = append(rpg.ArchivedFiles, *newArchivedFile)
	}
}

// Decrypts an integer from the RGSSADv3 archive.
//
// This function is meant for internal use by readRGSSAD.
//
// Returns the decrypted integer.
func (*RGSSADv3) decryptInteger(value int, key uint) int {
	result := int64(value) ^ int64(key)
	return int(result)
}

// Decrypts a filename from the RGSSADv3 archive.
//
// This function is meant for internal use by readRGSSAD.
//
// Returns the decrypted filename as a string.
func (*RGSSADv3) decryptFilename(encryptedName []byte, key uint) (decryptedName string) {
	keyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(keyBytes, uint32(key))
	i := 0
	j := 0

	for i < len(encryptedName) {
		j %= 4

		decryptedName += string(rune(encryptedName[i] ^ keyBytes[j]))

		i++
		j++
	}

	return decryptedName
}

// ExtractAllFiles
//
// See ExtractAllFiles in rgssad.go.
// A wrapper for ExtractAllFiles to remove the need for end-user casting to
// *RGSSAD.
func (rpg *RGSSADv3) ExtractAllFiles(outputDirectoryPath string, overrideExisting bool) error {
	return (*RGSSAD)(rpg).ExtractAllFiles(outputDirectoryPath, overrideExisting)
}
