package main

import (
	"fmt"
	"strings"

	"github.com/golang-fork-k/utils"

	"github.com/spf13/cobra"
)

func main() {

	var cmdGenerate = &cobra.Command{
		Use:   "generate [dir name]",
		Short: "Generate all folders",
		Long:  "Generate all sub folders",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			utils.Create(strings.Join(args, " "))
			if utils.IsCommandAvailable("kubectl") == false && utils.IsCommandAvailable("docker") == false {
				fmt.Println("please install kubectl first")
			}
		},
	}

	var cmdBuild = &cobra.Command{
		Use:   "build [tag]",
		Short: "build the docker image",
		Long:  "docker build from dockerfile",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			var tag string
			path = args[0]
			tag = args[1]
			utils.DockerBuild(path, tag, "Dockerfile")
			if utils.IsCommandAvailable("docker") == false {
				fmt.Println("please docker first")
			}
		},
	}

	var cmdDeploy = &cobra.Command{
		Use:   "deploy [tag]",
		Short: "deploy environment",
		Long:  "deploy kubernets ",
		// TODO: flag with environment
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var path string
			var name string
			path = args[0]
			name = args[1]
			utils.KubeDeploy(path, name)
			if utils.IsCommandAvailable("kubectl") == false {
				fmt.Println("please kubectl first")
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "k"}
	rootCmd.AddCommand(cmdGenerate, cmdBuild, cmdDeploy)
	rootCmd.Execute()

}
