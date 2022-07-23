package rpgo

import "testing"

func TestRPGMakerVersion(t *testing.T) {
	testArchive := "./test-rsrc/Game.rgssad"
	version := GetRPGMakerVersion(testArchive)

	if version != RPGMakerXp {
		t.Errorf("test/RPGMakerVersion: incorrect version; expected %d, got %d", RPGMakerXp, version)
	}

	testArchive = "./test-rsrc/Game.rgss2a"
	version = GetRPGMakerVersion(testArchive)

	if version != RPGMakerVx {
		t.Errorf("test/RPGMakerVersion: incorrect version; expected %d, got %d", RPGMakerVx, version)
	}

	testArchive = "./test-rsrc/Game.rgss3a"
	version = GetRPGMakerVersion(testArchive)

	if version != RPGMakerVxAce {
		t.Errorf("test/RPGMakerVersion: incorrect version; expected %d, got %d", RPGMakerVxAce, version)
	}

	testArchive = "./test-rsrc/invalid"
	version = GetRPGMakerVersion(testArchive)

	if version != RPGMakerInvalid {
		t.Errorf("test/RPGMakerVersion: incorrect version; expected %d, got %d", RPGMakerInvalid, version)
	}
}
