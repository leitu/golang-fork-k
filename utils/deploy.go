package utils

import (
	"log"
	"os/exec"
	"path/filepath"
)

func KubeDeploy(path, name string) {
	buildPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("Starting to deploy %s kubenerts service", name)
	serviceCmd := exec.Command("kubectl", "create", "-f", buildPath+"/infra/"+"kube/"+"service.yaml")
	log.Printf("%v", serviceCmd)
	if err := serviceCmd.Run(); err != nil {
		log.Printf("create kubernets service error")
	}

	log.Printf("Starting to deploy %s kubenerts configmap", name)
	configmapCmd := exec.Command("kubectl", "create", "-f", buildPath+"/infra/"+"kube/"+"configmap.yaml")
	log.Printf("%v", configmapCmd)
	if err := serviceCmd.Run(); err != nil {
		log.Printf("create kubernets service error")
	}

	log.Printf("Starting to deploy %s kubenerts deployment", name)
	deploymentCmd := exec.Command("kubectl", "create", "-f", buildPath+"/infra/"+"kube/"+"deployment.yaml")
	log.Printf("%v", deploymentCmd)
	if err := deploymentCmd.Run(); err != nil {
		log.Printf("create kubernets service error")
	}
}
