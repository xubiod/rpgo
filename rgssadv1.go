package rpgo

import "io"

//RPGMakerDecrypter.Decrypter/RGSSADv1.cs

type RGSSADv1 RGSSAD

func (RGSSADv1) Make(filepath string) *RGSSADv1 {
	created := new(RGSSADv1)
	created.Filepath = filepath

	version, err := ((*RGSSAD)(created)).GetVersion()

	if version != RGASSDv1 || err != nil {
		panic("MOMMMMM")
	}

	created.ReadRGSSAD()

	return created
}

func (rpg *RGSSADv1) ReadRGSSAD() {
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

		nameLen := DecryptInteger(num, &key)

		t = make([]byte, nameLen)

		rpg.ByteReader.Read(t)
		newArchivedFile.Name = DecryptFilename(t, &key)
		/// TODO: UNFUCK ALL OF THIS AND FINISH IT
	}
}

func DecryptInteger(value int, key *uint) int {
	result := value ^ int(*key)
	*key *= 7
	*key += 3

	return result
}

func DecryptFilename(encryptedName []byte, key *uint) string {
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
