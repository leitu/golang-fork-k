package utils

import (
	"os/exec"
)

func IsCommandAvailable(name string) bool {
	cmd := exec.Command("type", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
