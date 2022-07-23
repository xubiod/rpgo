package rpgo

import (
	"bytes"
	_ "crypto/md5"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractAllFiles_XP(t *testing.T) {
	extracted_answers := "./test-rsrc/extracted/xp/"

	var goat *RGSSADv1
	testArchive := "./test-rsrc/Game.rgssad"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ExtractAllFiles_XP: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv1(testArchive)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_XP: make rgssad err not nil:\n%s", err.Error())
		t.FailNow()
	}

	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/ExtractAllFiles_XP: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = goat.ExtractAllFiles(extractTo, true)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_XP: extract err not nil:\n%s", err.Error())
		t.FailNow()
	}

	// use os.SameFile(os.Stat(question), os.Stat(answer))

	all_questions, err := os.ReadDir(filepath.Join(extractTo, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_XP: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	all_answers, err := os.ReadDir(filepath.Join(extracted_answers, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_XP: readdir for answers err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(all_questions) != len(all_answers) {
		t.Errorf("test/ExtractAllFiles_XP: file count mismatch; expected %d, got %d", len(all_answers), len(all_questions))
		t.FailNow()
	}

	for idx := range all_questions {
		question_data, err := os.ReadFile(filepath.Join(extractTo, "Data", all_questions[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_XP: err getting question data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		answer_data, err := os.ReadFile(filepath.Join(extracted_answers, "Data", all_answers[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_XP: err getting answer data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		if !bytes.Equal(question_data, answer_data) {
			t.Errorf("test/ExtractAllFiles_XP: file data mismath at index %d (sorted alphabetically)", idx)
			t.FailNow()
		}
	}
}

func TestExtractAllFiles_VX(t *testing.T) {
	extracted_answers := "./test-rsrc/extracted/vx/"

	var goat *RGSSADv1
	testArchive := "./test-rsrc/Game.rgss2a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ExtractAllFiles_VX: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv1(testArchive)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VX: make rgssad err not nil:\n%s", err.Error())
		t.FailNow()
	}

	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VX: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = goat.ExtractAllFiles(extractTo, true)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VX: extract err not nil:\n%s", err.Error())
		t.FailNow()
	}

	// use os.SameFile(os.Stat(question), os.Stat(answer))

	all_questions, err := os.ReadDir(filepath.Join(extractTo, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VX: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	all_answers, err := os.ReadDir(filepath.Join(extracted_answers, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VX: readdir for answers err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(all_questions) != len(all_answers) {
		t.Errorf("test/ExtractAllFiles_VX: file count mismatch; expected %d, got %d", len(all_answers), len(all_questions))
		t.FailNow()
	}

	for idx := range all_questions {
		question_data, err := os.ReadFile(filepath.Join(extractTo, "Data", all_questions[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_VX: err getting question data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		answer_data, err := os.ReadFile(filepath.Join(extracted_answers, "Data", all_answers[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_VX: err getting answer data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		if !bytes.Equal(question_data, answer_data) {
			t.Errorf("test/ExtractAllFiles_VX: file data mismath at index %d (sorted alphabetically)", idx)
			t.FailNow()
		}
	}
}

func TestExtractAllFiles_VXAce(t *testing.T) {
	extracted_answers := "./test-rsrc/extracted/vxa/"

	var goat *RGSSADv3
	testArchive := "./test-rsrc/Game.rgss3a"

	if _, err := os.Stat(testArchive); errors.Is(err, os.ErrNotExist) {
		t.Error("test/ExtractAllFiles_VXAce: test config fail, archive doesn't exist")
		t.FailNow()
	}

	goat, err := MakeRGSSADv3(testArchive)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VXAce: make rgssad err not nil:\n%s", err.Error())
		t.FailNow()
	}

	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VXAce: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = goat.ExtractAllFiles(extractTo, true)

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VXAce: extract err not nil:\n%s", err.Error())
		t.FailNow()
	}

	// use os.SameFile(os.Stat(question), os.Stat(answer))

	all_questions, err := os.ReadDir(filepath.Join(extractTo, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VXAce: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	all_answers, err := os.ReadDir(filepath.Join(extracted_answers, "Data"))

	if err != nil {
		t.Errorf("test/ExtractAllFiles_VXAce: readdir for answers err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(all_questions) != len(all_answers) {
		t.Errorf("test/ExtractAllFiles_VXAce: file count mismatch; expected %d, got %d", len(all_answers), len(all_questions))
		t.FailNow()
	}

	for idx := range all_questions {
		question_data, err := os.ReadFile(filepath.Join(extractTo, "Data", all_questions[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_VXAce: err getting question data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		answer_data, err := os.ReadFile(filepath.Join(extracted_answers, "Data", all_answers[idx].Name()))

		if err != nil {
			t.Errorf("test/ExtractAllFiles_VXAce: err getting answer data at index %d (sorted alphabetically):\n%s", idx, err.Error())
			t.FailNow()
		}

		if !bytes.Equal(question_data, answer_data) {
			t.Errorf("test/ExtractAllFiles_VXAce: file data mismath at index %d (sorted alphabetically)", idx)
			t.FailNow()
		}
	}
}
