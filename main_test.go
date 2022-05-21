package main

import (
	"testing"
	"os"
)

func TestCreateFolder(t *testing.T) {
	dir := "test"
	path := "./test"
	os.RemoveAll(path)
	os.Mkdir(path, 0755)
	createFolder(dir, path)
	filesInDir, err := os.ReadDir(path)
	if err != nil {
		t.Error(err)
	}
	if len(filesInDir) != 1 {
		t.Error("Folder not created")
	}
	os.Remove(path + "/" + dir)
}


func TestCreateFolders(t *testing.T) {
	dir := "./test"
	createFolders(dir)
	filesInDir, err := os.ReadDir(dir)
	if err != nil {
		t.Error(err)
	}
	if len(filesInDir) != 9 {
		t.Error("Folders not created")
	}
	createFolder("Programs", dir)
	createFolder("Others", dir)
	createFolder("Folders", dir)
	filesInDir, err = os.ReadDir(dir)
	if err != nil {
		t.Error(err)
	}
	if len(filesInDir) > 9 {
		t.Error("Folders duplicated error")
	}
}


func TestMoveFile(t *testing.T) {
	os.RemoveAll("./test")
	os.Mkdir("./test", 0755)
	src := "./test/test.txt"
	dest := "./test/test2.txt"
	os.RemoveAll(dest)
	os.Create(src)
	moveFile(src, dest)
	filesInDir, err := os.ReadDir("./test")
	if err != nil {
		t.Error(err)
	}
	if len(filesInDir) != 1 {
		t.Error("File not moved")
	}
	os.Remove(dest)
}
