# RPGO

This is currently a direct translation of uuksuu's RPGMakerDecrypter project,
specifically the Decrypter namespace and the Tests namespace. You can go to that
version [here](https://github.com/uuksu/RPGMakerDecrypter/).

This was mainly an exercise with Go and the testing capabilities.

## Tests

**Current coverage as reported by `go test -cover`: 91.0%**

|File|Coverage|
|-|-|
|binary_utils.go|100%|
|project_generator.go|77.4%|
|rgssad.go|85.5%|
|rgssadv1.go|97.5%|
|rgssadv3.go|97.8%|
|rpgmaker_version.go|100.0%|

Currently all tests that were in the C# original have been implemented and all
pass.

Tests were added for encrypted file extracting for all supported versions. This
is redundant but all tests pass.
