package rpgo

import (
	"fmt"
	"io"
)

//RPGMakerDecrypter.Decrypter/RGSSADv1.cs

type RGSSADv1 RGSSAD

func (RGSSADv1) Make(filepath string) (*RGSSADv1, error) {
	created := new(RGSSADv1)
	created.Filepath = filepath

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv1 || err != nil {
		return nil, fmt.Errorf("rpgo/rgssadv3: version not v3 or this:\n%s", err.Error())
	}

	created.readRGSSAD()

	return created, nil
}

func (rpg *RGSSADv1) readRGSSAD() {
	key := uint(RGASSADv1Key)

	rpg.ByteReader.Seek(8, io.SeekStart)

	for {
		newArchivedFile := new(RPGMakerArchivedFile)

		t := make([]byte, 4)
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

func (*RGSSADv1) decryptInteger(value int, key *uint) int {
	result := value ^ int(*key)
	*key *= 7
	*key += 3

	return result
}

func (*RGSSADv1) decryptFilename(encryptedName []byte, key *uint) string {
	//decryptedName := make([]byte, len(encryptedName))
	var decryptedName string

	i := 0

	for i < len(encryptedName) {
		decryptedName += string(rune(encryptedName[i] ^ byte(*key&0xFF)))

		*key *= 7
		*key += 3

		i++
	}

	return decryptedName
}
