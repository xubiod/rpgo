package rpgo

import (
	"bytes"
	"io"
)

// RPGMakerDecrypter.Decrypter/BinaryUtils.cs

func ReadCString(reader *bytes.Reader, maxLength int) (string, error) {
	// startPostion := reader.
	stringLength, _ := reader.Seek(0, io.SeekCurrent)
	var readIn byte
	var err error
	var result string

	for stringLength < int64(maxLength) {
		readIn, err = reader.ReadByte()

		if readIn == 0 || err != nil {
			break
		}
		result += string(rune(readIn))
		stringLength++
	}

	return result, err
}
