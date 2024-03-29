package rpgo

// RPGMakerDecrypter.Decrypter/Constants.cs

const XpArchiveName = "Game.rgssad"    // Filename and extension of a typical RMXP encrypted archive.
const VxArchiveName = "Game.rgss2a"    // Filename and extension of a typical RMVX encrypted archive.
const VxAceArchiveName = "Game.rgss3a" // Filename and extension of a typical RMVXAce encrypted archive.

const XpProjectFileContent = "RPGXP 1.05"       // Contents of a typical RMXP project file generated by the latest version.
const VxProjectFileContent = "RPGVX 1.03"       // Contents of a typical RMVX project file generated by the latest version.
const VxAceProjectFileContent = "RPGVXAce 1.02" // Contents of a typical RMVXAce project file generated by the latest version.

const XpProjectFileExtension = "rxproj"     // The project file extension for a RMXP project.
const VxProjectFileExtension = "rvproj"     // The project file extension for a RMVX project.
const VxAceProjectFileExtension = "rvproj2" // The project file extension for a RMVXAce project.

const RGSSADHeader = "RGSSAD" // The header for a RPG Maker encrypted archive.

const RGASSDv1 = 1 // The version number for RMXP and RMVX archives.
const RGASSDv3 = 3 // The version number for RMVXAce archives.

const RGASSADv1Key uint = 0xDEADCAFE // The starting key for RMXP and RMVX archives.

// XpIniFileContents
//
// The contents of "Game.ini" for RMXP projects.
const XpIniFileContents = "[Game]\r\nLibrary=RGS104E.dll\r\nScripts=Data\\Scripts.rxdata\r\nTitle=DecryptedProject\r\nRTP1=Standard\r\nRTP2=\r\nRTP3="

// VxIniFileContents
//
// The contents of "Game.ini" for RMVX projects.
const VxIniFileContents = "[Game]\r\nRTP=RPGVX\r\nLibrary=RGS202E.dll\r\nScripts=Data\\Scripts.rvdata\r\nTitle=DecryptedProject"

// VxAceIniFileContents
//
// The contents of "Game.ini" for RMVXAce projects.
const VxAceIniFileContents = "[Game]\r\nRTP=RPGVXAce\r\nLibrary=RGSS300.dll\r\nScripts=Data\\Scripts.rvdata2\r\nTitle=DecryptedProject"
