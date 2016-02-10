package config

import (
	"os"
	"syscall"
)

func createOrOpenLockedFile(name string) (file *os.File, err error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
	} else {
		file, err = os.OpenFile(name, os.O_RDWR, 0644)
	}

	if err != nil {
		return
	}

	if flerr := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); flerr != nil {
		return file, flerr
	}

	return
}

func closeLockedFile(file *os.File) error {
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	return file.Close()
}
