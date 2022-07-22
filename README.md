# RPGO

This is currently a direct translation of uuksuu's RPGMakerDecrypter project,
specifically the Decrypter namespace and the Tests namespace. You can go to that
version [here](https://github.com/uuksu/RPGMakerDecrypter/).

This was mainly an exercise with Go and the testing capabilities.

## Tests

**Current coverage as reported by `go test -cover`: 75.1%** 

Currently all tests that were in the C# original have been implemented and all
pass.

Tests were added for encrypted file extracting for all supported versions. This
is redundant but all tests pass.

## Disclaimer

**The untested functions are not guaranteed to work, period. Do not use this as**
**a basis for any production-level project yet.**
