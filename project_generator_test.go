package rpgo

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateProject_XP(t *testing.T) {
	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/GenerateProject_XP: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = GenerateProject(RPGMakerXp, extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_XP: generateproject err not nil:\n%s", err.Error())
		t.FailNow()
	}

	directory, err := os.ReadDir(extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_XP: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(directory) != 2 {
		t.Errorf("test/GenerateProject_XP: file count mismatch; expected 2, got %d", len(directory))
		t.FailNow()
	}

	if directory[0].Name() != "Game.ini" {
		t.Errorf("test/GenerateProject_XP: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if directory[1].Name() != fmt.Sprintf("Game.%s", XpProjectFileExtension) {
		t.Errorf("test/GenerateProject_XP: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}
}

func TestGenerateProject_VX(t *testing.T) {
	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/GenerateProject_VX: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = GenerateProject(RPGMakerVx, extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_VX: generateproject err not nil:\n%s", err.Error())
		t.FailNow()
	}

	directory, err := os.ReadDir(extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_VX: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(directory) != 2 {
		t.Errorf("test/GenerateProject_VX: file count mismatch; expected 2, got %d", len(directory))
		t.FailNow()
	}

	if directory[0].Name() != "Game.ini" {
		t.Errorf("test/GenerateProject_VX: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if directory[1].Name() != fmt.Sprintf("Game.%s", VxProjectFileExtension) {
		t.Errorf("test/GenerateProject_VX: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}
}

func TestGenerateProject_VXAce(t *testing.T) {
	extractTo, err := os.MkdirTemp(os.TempDir(), "rpgo-testsuite-*")

	if err != nil {
		t.Errorf("test/GenerateProject_VXAce: mkdir temp err not nil:\n%s", err.Error())
		t.FailNow()
	}

	err = GenerateProject(RPGMakerVxAce, extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_VXAce: generateproject err not nil:\n%s", err.Error())
		t.FailNow()
	}

	directory, err := os.ReadDir(extractTo)

	if err != nil {
		t.Errorf("test/GenerateProject_VXAce: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if len(directory) != 2 {
		t.Errorf("test/GenerateProject_VXAce: file count mismatch; expected 2, got %d", len(directory))
		t.FailNow()
	}

	if directory[0].Name() != "Game.ini" {
		t.Errorf("test/GenerateProject_VXAce: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}

	if directory[1].Name() != fmt.Sprintf("Game.%s", VxAceProjectFileExtension) {
		t.Errorf("test/GenerateProject_VXAce: readdir for questions err not nil:\n%s", err.Error())
		t.FailNow()
	}
}
