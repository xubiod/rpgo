package rpgo

import (
	"bytes"
	"io"
)

// RPGMakerDecrypter.Decrypter/BinaryUtils.cs

// Reads a C string at the current position of the bytes.Reader, with nil as
// error.
//
// If the string is longer than the maximum length, it will return what it read
// (length of returned string is <= maxLength).
//
// Errors will make it return what it read with an error.
func ReadCString(reader *bytes.Reader, maxLength int) (string, error) {
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
