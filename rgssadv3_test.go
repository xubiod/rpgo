package rpgo

import (
	"errors"
	"os"
	"testing"
)

func TestRGSSADv3_VXAce(t *testing.T) {
	var goat *RGSSADv3
	testArchive := "./test-rsrc/Game.rgss3a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/RGSSADv3_VXAce: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv3(testArchive)

	if err != nil {
		t.Errorf("test/RGSSADv3_VXAce: err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(goat.ArchivedFiles) != 16 {
		t.Errorf("test/RGSSADv3_VXAce: archived files length not expected; wanted 16, got %d", len(goat.ArchivedFiles))
		t.Fail()
	}

	expectedNames := []string{"Data\\Actors.rvdata2", "Data\\Animations.rvdata2", "Data\\Armors.rvdata2"}
	expectedOffsets := []int64{605, 3637, 222096}
	expectedSizes := []int{3032, 218459, 11472}
	expectedKeys := []uint{0x00000029, 0x00004823, 0x000018BE}

	for idx := range expectedNames {
		if goat.ArchivedFiles[idx].Name != expectedNames[idx] {
			t.Errorf("test/RGSSADv3_VXAce: archived files name not expected for index %d; wanted %s, got %s", idx, expectedNames[idx], goat.ArchivedFiles[idx].Name)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Offset != expectedOffsets[idx] {
			t.Errorf("test/RGSSADv3_VXAce: archived files offset not expected for index %d; wanted %d, got %d", idx, expectedOffsets[idx], goat.ArchivedFiles[idx].Offset)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Size != expectedSizes[idx] {
			t.Errorf("test/RGSSADv3_VXAce: archived files size not expected for index %d; wanted %d, got %d", idx, expectedSizes[idx], goat.ArchivedFiles[idx].Size)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Key != expectedKeys[idx] {
			t.Errorf("test/RGSSADv3_VXAce: archived files key not expected for index %d; wanted %X, got %X", idx, expectedKeys[idx], goat.ArchivedFiles[idx].Key)
			t.Fail()
		}

		if t.Failed() {
			t.FailNow()
		}
	}
}
