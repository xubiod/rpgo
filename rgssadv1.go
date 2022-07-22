package rpgo

import (
	"encoding/binary"
	"fmt"
	"io"
)

//RPGMakerDecrypter.Decrypter/RGSSADv1.cs

type RGSSADv1 RGSSAD

func MakeRGSSADv1(filepath string) (*RGSSADv1, error) {
	created := (*RGSSADv1)(MakeRGSSAD(filepath))

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv1 || err != nil {
		return nil, fmt.Errorf("rpgo/rgssadv1: version not v1, or err'd to this:\n%s", err.Error())
	}

	created.readRGSSAD()

	return created, nil
}

func (rpg *RGSSADv1) readRGSSAD() {
	key := RGASSADv1Key

	t := make([]byte, 4)

	rpg.ByteReader.Seek(8, io.SeekStart)

	for {
		newArchivedFile := new(RPGMakerArchivedFile)

		rpg.ByteReader.Read(t)
		num := int(binary.LittleEndian.Uint32(t))

		nameLen := rpg.decryptInteger(int(num), &key)

		u := make([]byte, nameLen)

		rpg.ByteReader.Read(u)
		newArchivedFile.Name = rpg.decryptFilename(u, &key)

		rpg.ByteReader.Read(t)
		num = int(binary.LittleEndian.Uint32(t))

		newArchivedFile.Size = rpg.decryptInteger(num, &key)

		newArchivedFile.Offset, _ = rpg.ByteReader.Seek(0, io.SeekCurrent)
		newArchivedFile.Key = key

		rpg.ArchivedFiles = append(rpg.ArchivedFiles, *newArchivedFile)

		rpg.ByteReader.Seek(int64(newArchivedFile.Size), io.SeekCurrent)

		// status, _ := rpg.ByteReader.Seek(0, io.SeekCurrent)
		if rpg.ByteReader.Len() == 0 { // len returns UNREAD bytes what the FUCK
			break
		}
	}
}

func (*RGSSADv1) decryptInteger(value int, key *uint) int {
	result := int64(value) ^ int64(*key)
	*key *= 7
	*key += 3
	*key &= 0xFFFFFFFF

	return int(result)
}

func (*RGSSADv1) decryptFilename(encryptedName []byte, key *uint) string {
	//decryptedName := make([]byte, len(encryptedName))
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
