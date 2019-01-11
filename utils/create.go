package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	TestfileName       = "test.yaml"
	IgnorefileName     = ".testingore"
	DeploymentfileName = "deployment.yaml"
	ServicefileName    = "service.yaml"
	ConfigmapfileName  = "configmap.yaml"
	DockerfileValue    = "Dockerfile"
)

var SubFolderName = []string{"infra", "app", "bin", "config", "db", "lib", "log", "spec", "tmp", "vendor"}
var InfraSubFolderName = []string{"ci", "docs", "keys", "kube", "secrets"}

func CreateDirIfNotExist(dir string) {
	//fmt.Printf("%s", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("%s", err)
		}
	}
}

// Dir means the top folder
func Create(dir string) {

	path, err := filepath.Abs(dir)
	if err != nil {
		log.Printf("%s", err)
	}
	if fi, err := os.Stat(path); err != nil {
		log.Printf("%s", err)
	} else if !fi.IsDir() {
		log.Printf("no such directory %s", path)
	}

	for _, v := range SubFolderName {
		sdir := filepath.Join(path, v)
		if v == "infra" {
			for _, vi := range InfraSubFolderName {
				cidir := filepath.Join(sdir, vi)
				CreateDirIfNotExist(cidir)
			}
		}
		CreateDirIfNotExist(sdir)
	}

	files := []struct {
		path    string
		content []byte
	}{
		{
			//test.yaml
			path:    filepath.Join(path, TestfileName),
			content: []byte(fmt.Sprintf(defaultValues, TestfileName)),
		},
		{
			//Dockerfile
			path:    filepath.Join(path, DockerfileValue),
			content: []byte(defaultDockerValue),
		},
		{
			// .helmignore
			path:    filepath.Join(path, IgnorefileName),
			content: []byte(defaultIgnore),
		},
		{
			//deployment.yaml
			path:    filepath.Join(path, SubFolderName[0], InfraSubFolderName[3], DeploymentfileName),
			content: []byte(fmt.Sprintf(defaultDeploymentVaule, "bot-xyz")),
		},
		{
			//service.yaml
			path:    filepath.Join(path, SubFolderName[0], InfraSubFolderName[3], ServicefileName),
			content: []byte(fmt.Sprintf(defaultServiceVaule, "bot-xyz")),
		},
		{
			//service.yaml
			path:    filepath.Join(path, SubFolderName[0], InfraSubFolderName[3], ConfigmapfileName),
			content: []byte(fmt.Sprintf(defaultConfigmapValue, "bot-xyz")),
		},
	}

	for _, file := range files {
		if _, err := os.Stat(file.path); err == nil {
			// File exists and is okay. Skip it.
			continue
		}
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			log.Printf("%s", err)
		}
	}

}
