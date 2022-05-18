package pkg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	ErrFileExist    = errors.New("file exist")
	ErrFileNotExist = errors.New("file not exist")
)

func git(args ...string) *exec.Cmd {
	return exec.Command("git", args...)
}

func isExist(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Set(file, cmd string) error {
	info, err := os.Stat(filepath.Dir(file))
	if err != nil {
		color.Red(`can't create hook, %s directory doesn't exist (try running shiba install)`, file)
		return ErrFileNotExist
	}
	if !isExist(info.Name()) {
		return ErrFileNotExist
	}
	f, err := os.Create(file)
	if err != nil {
		return ErrFileNotExist
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf(
		`#!/usr/bin/env sh
. "$(dirname -- "$0")/_/shiba.sh"
%s
`, cmd))
	color.Green("Created %s", file)
	return nil
}

func Add(file, cmd string) error {
	if isExist(file) {
		f, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
		f.WriteString(fmt.Sprintf("%s\n", cmd))
		f.Close()
		color.Green("Updated %s", file)
	} else {
		return Set(file, cmd)
	}
	return nil
}

func Install() error {
	return nil
}

func Uninstall() error {
	git("config", "--unset", "core.hooksPath")
	return nil
}
