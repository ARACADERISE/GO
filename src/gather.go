package src

import (
	"os"
	"log"
)

type DirInfo struct {
	path		string
	file		string
	full_path	string
	valid		bool
}

func Gather_default() *DirInfo {
	curr_dir, err := os.Getwd()

	if err != nil {
		log.Fatal("Error finding current path")
	}

	info := &DirInfo { path: curr_dir, file: "main.go", full_path: curr_dir + "/main.go", valid: false } // valid is false for now, lets see if we can change that!

	_, _err := os.Stat(info.full_path)

	if _err != nil {
		log.Fatal(_err)
	}

	info.valid = true

	return info
}

func Gather_default_file(filename string) *DirInfo {
	curr_dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	info := &DirInfo{ path: curr_dir, file: filename, full_path: curr_dir + "/" + filename, valid: false } // valid is false for now, lets see if we can change that!

	_, _err := os.Stat(info.full_path)

	if _err != nil {
		log.Fatal(_err)
	}

	info.valid = true

	return info
}
