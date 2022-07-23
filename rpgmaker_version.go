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
