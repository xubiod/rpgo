package rpgo

import (
	"bytes"
)

// RPGMakerDecrypter.Decrypter/BinaryUtils.cs

func ReadCString(reader bytes.Reader, maxLength int) (string, error) {
	// startPostion := reader.
	stringLength := 0
	var readIn byte
	var err error
	var result string

	for stringLength < maxLength {
		readIn, err = reader.ReadByte()

		if readIn == 0 || err != nil {
			break
		}
		result += string(rune(readIn))
		stringLength++
	}

	return result, err
}
