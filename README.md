# RPGO [![Go](https://github.com/xubiod/rpgo/actions/workflows/go.yml/badge.svg)](https://github.com/xubiod/rpgo/actions/workflows/go.yml)

This is currently a direct translation of uuksuu's RPGMakerDecrypter project,
specifically the Decrypter namespace and the Tests namespace. You can go to that
version [here](https://github.com/uuksu/RPGMakerDecrypter/).

This was mainly an exercise with Go and the testing capabilities.

## Tests

**Current coverage as reported by `go test -cover`: 90.9%**

|File|Coverage|
|-|-|
|[binary_utils.go](binary_utils.go)|100%|
|[project_generator.go](project_generator.go)|77.4%|
|[rgssad.go](rgssad.go)|85.5%|
|[rgssadv1.go](rgssadv1.go)|97.6%|
|[rgssadv3.go](rgssadv3.go)|97.9%|
|[rpgmaker_version.go](rpgmaker_version.go)|100.0%|
