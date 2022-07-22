package rpgo

import (
	"errors"
	"os"
	"testing"
)

func TestRGSSADv1_XP(t *testing.T) {
	var goat *RGSSADv1
	testArchive := "./test-rsrc/Game.rgssad"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/RGSSADv1_XP: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv1(testArchive)

	if err != nil {
		t.Errorf("test/RGSSADv1_XP: err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(goat.ArchivedFiles) != 16 {
		t.Errorf("test/RGSSADv1_XP: archived files length not expected; wanted 16, got %d", len(goat.ArchivedFiles))
		t.Fail()
	}

	expectedNames := []string{"Data\\Actors.rxdata", "Data\\Animations.rxdata", "Data\\Armors.rxdata"}
	expectedOffsets := []int64{34, 11045, 147314}
	expectedSizes := []int{10981, 136243, 4285}
	expectedKeys := []uint{0x7B7448AE, 0x366D564E, 0x222699FE}

	for idx := range expectedNames {
		if goat.ArchivedFiles[idx].Name != expectedNames[idx] {
			t.Errorf("test/RGSSADv1_XP: archived files name not expected for index %d; wanted %s, got %s", idx, expectedNames[idx], goat.ArchivedFiles[idx].Name)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Offset != expectedOffsets[idx] {
			t.Errorf("test/RGSSADv1_XP: archived files offset not expected for index %d; wanted %d, got %d", idx, expectedOffsets[idx], goat.ArchivedFiles[idx].Offset)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Size != expectedSizes[idx] {
			t.Errorf("test/RGSSADv1_XP: archived files size not expected for index %d; wanted %d, got %d", idx, expectedSizes[idx], goat.ArchivedFiles[idx].Size)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Key != expectedKeys[idx] {
			t.Errorf("test/RGSSADv1_XP: archived files key not expected for index %d; wanted %X, got %X", idx, expectedKeys[idx], goat.ArchivedFiles[idx].Key)
			t.Fail()
		}

		if t.Failed() {
			t.FailNow()
		}
	}
}

func TestRGSSADv1_VX(t *testing.T) {
	var goat *RGSSADv1
	testArchive := "./test-rsrc/Game.rgss2a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/RGSSADv1_VX: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv1(testArchive)

	if err != nil {
		t.Errorf("test/RGSSADv1_VX: err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(goat.ArchivedFiles) != 16 {
		t.Errorf("test/RGSSADv1_VX: archived files length not expected; wanted 16, got %d", len(goat.ArchivedFiles))
		t.Fail()
	}

	expectedNames := []string{"Data\\Actors.rvdata", "Data\\Animations.rvdata", "Data\\Areas.rvdata"}
	expectedOffsets := []int64{34, 10951, 139280}
	expectedSizes := []int{10887, 128304, 4}
	expectedKeys := []uint{0x7B7448AE, 0x366D564E, 0x04E0F16D}

	for idx := range expectedNames {
		if goat.ArchivedFiles[idx].Name != expectedNames[idx] {
			t.Errorf("test/RGSSADv1_VX: archived files name not expected for index %d; wanted %s, got %s", idx, expectedNames[idx], goat.ArchivedFiles[idx].Name)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Offset != expectedOffsets[idx] {
			t.Errorf("test/RGSSADv1_VX: archived files offset not expected for index %d; wanted %d, got %d", idx, expectedOffsets[idx], goat.ArchivedFiles[idx].Offset)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Size != expectedSizes[idx] {
			t.Errorf("test/RGSSADv1_VX: archived files size not expected for index %d; wanted %d, got %d", idx, expectedSizes[idx], goat.ArchivedFiles[idx].Size)
			t.Fail()
		}

		if goat.ArchivedFiles[idx].Key != expectedKeys[idx] {
			t.Errorf("test/RGSSADv1_VX: archived files key not expected for index %d; wanted %X, got %X", idx, expectedKeys[idx], goat.ArchivedFiles[idx].Key)
			t.Fail()
		}

		if t.Failed() {
			t.FailNow()
		}
	}
}
