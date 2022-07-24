# RPGO-CLI

This is a command-line tool that is close to basic feature parity with
**RPG-Patcher**. You can see what it can do under [Actions](#actions) or running
it without any flags.

You can run this CLI with Go 1.18.4 (currently tested version):
> `go run rpgo-cli.go`.

This file is meant as a more extensive and exhaustive version of the default
behaviour (no flags/incorrect flags).

- [RPGO-CLI](#rpgo-cli)
  - [Actions](#actions)
    - [Extracting files](#extracting-files)
    - [Lists files in archives](#lists-files-in-archives)
    - [Only generating project files](#only-generating-project-files)

## Actions

### Extracting files

**Action names: `extract`, `decrypt`, `dump`**

**Flags used: `i`, `o`, `overwrite-files`**

Extracts all files from the encrypted archive given via `-i`. All
files are extracted into the directory given via `-o`. If
`-overwrite-files` is present with t, true or 1, files will override existing
files; otherwise it ignores the file and continues.

### Lists files in archives

**Action names: `files`, `list`, `ls`, `dir`**

**Flags used: `o`**

Reads an encrypted archive and prints out their names and sizes into standard
output. You can configure to display megabytes, kilobytes, bytes by setting `-o`
to `kilo`; mebibytes, kibibytes, bytes by setting `-o` to `kibi`; or just bytes
by setting `-o` to `bytes`.

The output format is as follows:

`Name [tab] (Size [tab] Unit)`

Name is the file name with its directory in the archive.

Size is either an integer for bytes or floats for other units.

Unit is `B`, `KB`, `MB`, `GB` for `-o kilo`; `B`, `KiB`, `MiB`, `GiB` for
`-o kibi`; `bytes` for `-o bytes`.

### Only generating project files

**Action names: `justproject`**

**Flags used: `i`, `o`**

Just generates a project files for the RPG Maker version given via `-i`. The
generated files are placed in the output directory given via `-o`.

Acceptable input values are:

- For XP, `XP`, `RPG Maker XP`, `xp`, `rmxp`, `rgssad`, `Game.rgssad`, `rxproj`
- For VX, `VX`, `RPG Maker VX`, `vx`, `rmvx`, `rgss2a`, `Game.rgss2a`, `rvproj`
- For VX Ace, `VXAce`, `VX Ace`, `RPG Maker VXAce`, `RPG Maker VX Ace`, `vxace`,
  `rmvxace`, `rmvxa`, `rgss3a`, `Game.rgss3a`, `rvproj2`
