package rpgo

// RPGMakerDecrypter.Decrypter/Constants.cs

const XpArchiveName = "Game.rgssad"
const VxArchiveName = "Game.rgss2a"
const VxAceArchiveName = "Game.rgss3a"

const XpProjectFileContent = "RPGXP 1.05"
const VxProjectFileContent = "RPGVX 1.03"
const VxAceProjectFileContent = "RPGVXAce 1.02"

const XpProjectFileExtension = "rxproj"
const VxProjectFileExtension = "rvproj"
const VxAceProjectFileExtension = "rvproj2"

const RGSSADHeader = "RGSSAD"

const RGASSDv1 = 1
const RGASSDv3 = 3

const RGASSADv1Key uint = 0xDEADCAFE

const XpIniFileContents = "[Game]\r\nLibrary=RGS104E.dll\r\nScripts=Data\\Scripts.rxdata\r\nTitle=DecryptedProject\r\nRTP1=Standard\r\nRTP2=\r\nRTP3="
const VxIniFileContents = "[Game]\r\nRTP=RPGVX\r\nLibrary=RGS202E.dll\r\nScripts=Data\\Scripts.rvdata\r\nTitle=DecryptedProject"
const VxAceIniFileContents = "[Game]\r\nRTP=RPGVXAce\r\nLibrary=RGSS300.dll\r\nScripts=Data\\Scripts.rvdata2\r\nTitle=DecryptedProject"
