package rpgo

import (
	"encoding/binary"
	"fmt"
	"io"
)

//RPGMakerDecrypter.Decrypter/RGSSADv1.cs

type RGSSADv1 RGSSAD

// Creates a new RGSSADv1 structure and configures it for use.
//
// Returns a pointer to the created structure, nil on success, nil and error
// otherwise.
func NewRGSSADv1(filepath string) (*RGSSADv1, error) {
	created := (*RGSSADv1)(NewRGSSAD(filepath))

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv1 || err != nil {
		return nil, fmt.Errorf("rpgo/rgssadv1: version not v1, or err'd to this:\n%s", err.Error())
	}

	created.readRGSSAD()

	return created, nil
}

// Reads the encrypted RGSSADv1 archive and generates a ArchivedFile slice for
// the RGSSAD structure.
//
// This function is meant for internal use in NewRGSSADv1.
func (rpg *RGSSADv1) readRGSSAD() {
	key := RGASSADv1Key

	t := make([]byte, 4)

	rpg.ByteReader.Seek(8, io.SeekStart)

	for {
		newArchivedFile := new(RPGMakerArchivedFile)

		// NAME
		rpg.ByteReader.Read(t)
		num := int(binary.LittleEndian.Uint32(t))

		nameLen := rpg.decryptInteger(int(num), &key)

		u := make([]byte, nameLen)

		rpg.ByteReader.Read(u)
		newArchivedFile.Name = rpg.decryptFilename(u, &key)

		// SIZE
		rpg.ByteReader.Read(t)
		num = int(binary.LittleEndian.Uint32(t))

		newArchivedFile.Size = rpg.decryptInteger(num, &key)

		// OFFSET, KEY
		newArchivedFile.Offset, _ = rpg.ByteReader.Seek(0, io.SeekCurrent)
		newArchivedFile.Key = key

		rpg.ArchivedFiles = append(rpg.ArchivedFiles, *newArchivedFile)

		rpg.ByteReader.Seek(int64(newArchivedFile.Size), io.SeekCurrent)

		if rpg.ByteReader.Len() == 0 {
			break
		}
	}
}

// Decrypts an integer from the RGSSADv1 archive, modifying the key for the next
// decryption afterwards.
//
// This function is meant for internal use by readRGSSAD.
//
// Returns the decrypted integer.
func (*RGSSADv1) decryptInteger(value int, key *uint) int {
	result := int64(value) ^ int64(*key)
	*key *= 7
	*key += 3
	*key &= 0xFFFFFFFF

	return int(result)
}

// Decrypts a filename from the RGSSADv1 archive, modifying the key for the next
// decryption afterwards.
//
// This function is meant for internal use by readRGSSAD.
//
// Returns the decrypted filename as a string.
func (*RGSSADv1) decryptFilename(encryptedName []byte, key *uint) string {
	var decryptedName string

	i := 0

	for i < len(encryptedName) {
		decryptedName += string(rune(encryptedName[i] ^ byte(*key&0xFF)))

		*key *= 7
		*key += 3
		*key &= 0xFFFFFFFF

		i++
	}

	return decryptedName
}

// See ExtractAllFiles in rgssad.go
//
// A wrapper for ExtractAllFiles to remove the need for end-user casting to
// *RGSSAD.
func (rpg *RGSSADv1) ExtractAllFiles(outputDirectoryPath string, overrideExisting bool) error {
	return (*RGSSAD)(rpg).ExtractAllFiles(outputDirectoryPath, overrideExisting)
}
