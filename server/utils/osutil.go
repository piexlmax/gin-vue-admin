package utils

import (
	"os"
	"syscall"
)

// Will return true if the process with PID exists.
func CheckPid(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}
