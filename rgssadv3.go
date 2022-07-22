package rpgo

import (
	"encoding/binary"
	"fmt"
	"io"
)

//RPGMakerDecrypter.Decrypter/RGSSADv3.cs

type RGSSADv3 RGSSAD

func (RGSSADv3) Make(filepath string) (*RGSSADv3, error) {
	created := new(RGSSADv3)
	created.Filepath = filepath

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv1 || err != nil {
		return nil, fmt.Errorf("rpgo/rgssadv3: version not v3 or this:\n%s", err.Error())
	}

	created.readRGSSAD()

	return created, nil
}

func (rpg *RGSSADv3) readRGSSAD() {
	rpg.ByteReader.Seek(8, io.SeekStart)

	t := make([]byte, 4)
	rpg.ByteReader.Read(t)
	var num int
	for _, e := range t {
		num = num << 8
		num |= int(e)
	}

	key := uint(num)

	key *= 9
	key += 3

	// key := uint(RGASSADv1Key)

	for {
		newArchivedFile := new(RPGMakerArchivedFile)

		t = make([]byte, 4)
		rpg.ByteReader.Read(t)
		var num int
		for _, e := range t {
			num = num << 8
			num |= int(e)
		}

		nameLen := rpg.decryptInteger(num, &key)

		t = make([]byte, nameLen)

		rpg.ByteReader.Read(t)
		newArchivedFile.Name = rpg.decryptFilename(t, &key)

		// SIZE

		t = make([]byte, 4)
		rpg.ByteReader.Read(t)
		num = 0
		for _, e := range t {
			num = num << 8
			num |= int(e)
		}

		newArchivedFile.Size = rpg.decryptInteger(num, &key)

		// END SIZE

		newArchivedFile.Offset, _ = rpg.ByteReader.Seek(0, io.SeekCurrent)
		newArchivedFile.Key = key

		rpg.ArchivedFiles = append(rpg.ArchivedFiles, *newArchivedFile)
		/// TODO: UNFUCK ALL OF THIS

		status, _ := rpg.ByteReader.Seek(0, io.SeekCurrent)
		if status == int64(rpg.ByteReader.Len()) {
			break
		}
	}
}

func (*RGSSADv3) decryptInteger(value int, key *uint) int {
	result := uint(value) ^ *key
	// *key *= 7
	// *key += 3

	return int(result)
}

func (*RGSSADv3) decryptFilename(encryptedName []byte, key *uint) string {
	var decryptedName string

	keyBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(keyBytes, uint32(*key))
	i := 0
	j := 0

	for i < len(encryptedName) {
		// if j == 4 {
		// 	j = 0
		// }
		j %= 4

		decryptedName += string(rune(encryptedName[i] ^ keyBytes[j]))

		i++
		j++
	}

	return decryptedName
}
