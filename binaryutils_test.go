package rpgo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestArchiveHeader_XP(t *testing.T) {
	testArchive := "./test-rsrc/Game.rgssad"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ArchiveHeader_XP: test config fail, archive doesn't exist")
		t.FailNow()
	}

	f, _ := os.Open(testArchive)

	data, _ := ioutil.ReadAll(f)
	byteReader := *bytes.NewReader(data)
	f.Close()

	header, err := ReadCString(&byteReader, 7)

	if err != nil {
		t.Errorf("test/ArchiveHeader_XP: ReadCString error:\n%s", err)
		t.Fail()
	}

	if header != RGSSADHeader {
		t.Errorf("test/ArchiveHeader_XP: header incorrect; expected %s, got %s", RGSSADHeader, header)
	}

	version, _ := byteReader.ReadByte()

	if version != 1 {
		t.Errorf("test/ArchiveHeader_XP: version mismatch; expected 1, got %d", version)
		t.FailNow()
	}

}

func TestArchiveHeader_VX(t *testing.T) {
	testArchive := "./test-rsrc/Game.rgss2a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ArchiveHeader_VX: test config fail, archive doesn't exist")
		t.FailNow()
	}

	f, _ := os.Open(testArchive)

	data, _ := ioutil.ReadAll(f)
	byteReader := *bytes.NewReader(data)
	f.Close()

	header, err := ReadCString(&byteReader, 7)

	if err != nil {
		t.Errorf("test/ArchiveHeader_VX: ReadCString error:\n%s", err)
		t.Fail()
	}

	if header != RGSSADHeader {
		t.Errorf("test/ArchiveHeader_XP: header incorrect; expected %s, got %s", RGSSADHeader, header)
	}

	version, _ := byteReader.ReadByte()

	if version != 1 {
		t.Errorf("test/ArchiveHeader_VX: version mismatch; expected 1, got %d", version)
		t.FailNow()
	}
}

func TestArchiveHeader_VXAce(t *testing.T) {
	testArchive := "./test-rsrc/Game.rgss3a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ArchiveHeader_VXAce: test config fail, archive doesn't exist")
		t.FailNow()
	}

	f, _ := os.Open(testArchive)

	data, _ := ioutil.ReadAll(f)
	byteReader := *bytes.NewReader(data)
	f.Close()

	header, err := ReadCString(&byteReader, 7)

	if err != nil {
		t.Errorf("test/ArchiveHeader_VXAce: ReadCString error:\n%s", err)
		t.Fail()
	}

	if header != RGSSADHeader {
		t.Errorf("test/ArchiveHeader_XP: header incorrect; expected %s, got %s", RGSSADHeader, header)
	}

	version, _ := byteReader.ReadByte()

	if version != 3 {
		t.Errorf("test/ArchiveHeader_VXAce: version mismatch; expected 3, got %d", version)
		t.FailNow()
	}
}
