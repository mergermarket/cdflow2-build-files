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
	actualPath, err := app.SaveData(targetDirectory, testFile)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	actualFilename := path.Base(actualPath)
	if actualFilename != "test.txt" {
		t.Fatalf("got '%v' filename, expected '%v'", actualFilename, "test.txt")
	}
	rawContents, err := ioutil.ReadFile(path.Join(targetDirectory, actualPath))
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
	actualPath, err := app.SaveData(targetDirectory, testDir)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	actualDirname := path.Base(actualPath)
	if actualDirname != "test" {
		t.Fatalf("got '%v' dirname, expected '%v'", actualDirname, "test")
	}
	rawContents, err := ioutil.ReadFile(path.Join(targetDirectory, actualPath, "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	contents := string(rawContents)
	if contents != "saved file" {
		t.Fatalf("got '%v', expected '%v'", contents, "saved file")
	}
}
