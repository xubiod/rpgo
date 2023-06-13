package rpgo

import "path/filepath"

// RPGMakerDecrypter.Decrypter/RPGMakerVersion.cs

type RPGMakerVersion int

const (
	RPGMakerInvalid RPGMakerVersion = iota + 1
	RPGMakerXp
	RPGMakerVx
	RPGMakerVxAce
)

// GetRPGMakerVersion
//
// Gets an input filepath to an encrypted archive as a string, returns what
// RPGMakerVersion it is.
//
// Returns RPGMakerInvalid for an invalid archive file name and extension.
func GetRPGMakerVersion(inputPath string) RPGMakerVersion {
	switch filepath.Base(inputPath) {
	case XpArchiveName:
		return RPGMakerXp
	case VxArchiveName:
		return RPGMakerVx
	case VxAceArchiveName:
		return RPGMakerVxAce
	default:
		return RPGMakerInvalid
	}
}
