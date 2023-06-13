package rpgo

// RPGMakerDecrypter.Decrypter/ArchivedFile.cs

// RPGMakerArchivedFile
//
// A file in an RPG Maker encrypted archive.
type RPGMakerArchivedFile struct {
	Name   string // The file name and path in the archive
	Size   int    // The size of the file in bytes
	Offset int64  // The offset in the archive where the data is in bytes
	Key    uint   // The key to decrypt the file contents
}
