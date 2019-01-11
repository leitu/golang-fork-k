package utils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func IsDockerfileAvailable(dockerfile string) bool {
	path, err := filepath.Abs(dockerfile)
	if err != nil {
		log.Printf("%s", err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		log.Printf("no such directory %s", path)
		return true
	}
	return false
}

func DockerBuild(path, tag, dockerfile string) {
	buildPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("%s", err)
	}
	if !IsDockerfileAvailable(dockerfile) {
		log.Printf("Starting to build docker image")
		cmd := exec.Command("docker", "build", "-t", tag, "-f", buildPath+"/"+dockerfile, buildPath)
		log.Printf("%v", cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Build error")
		}
	}
}
