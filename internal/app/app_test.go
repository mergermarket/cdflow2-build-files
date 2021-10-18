package app_test

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/mergermarket/cdflow2-build-files/internal/app"
)

func TestSaveFile(t *testing.T) {
	// Given
	targetDirectory, err := ioutil.TempDir("", "cdflow2-build-files-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(targetDirectory)
	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Join(path.Dir(filename), "test", "test.txt")

	// When
	savedPath, err := app.SaveData(targetDirectory, testFile)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	savedFilename := path.Base(savedPath)
	if savedFilename != "test.txt" {
		t.Fatalf("got '%v' filename, expected '%v'", savedFilename, "test.txt")
	}
	rawContents, err := ioutil.ReadFile(savedPath)
	if err != nil {
		t.Fatal(err)
	}
	contents := string(rawContents)
	if contents != "saved file" {
		t.Fatalf("got '%v', expected '%v'", contents, "saved file")
	}
}

func TestSaveDir(t *testing.T) {
	// Given
	targetDirectory, err := ioutil.TempDir("", "cdflow2-build-files-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(targetDirectory)
	_, filename, _, _ := runtime.Caller(0)
	testDir := path.Join(path.Dir(filename), "test")

	// When
	savedPath, err := app.SaveData(targetDirectory, testDir)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	savedDirname := path.Base(savedPath)
	if savedDirname != "test" {
		t.Fatalf("got '%v' dirname, expected '%v'", savedDirname, "test")
	}
	rawContents, err := ioutil.ReadFile(path.Join(savedPath, "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	contents := string(rawContents)
	if contents != "saved file" {
		t.Fatalf("got '%v', expected '%v'", contents, "saved file")
	}
}
