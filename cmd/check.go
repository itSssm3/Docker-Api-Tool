package cmd

import (
	"context"
	"docker-api/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check DockerApi connection",
	Run:   Check,
}

func init() {
	checkCmd.Flags().StringVar(&address, "address", "", "Target API opened IP and port: tcp://127.0.0.1:2375")
	checkCmd.Flags().StringVar(&clientversion, "clientversion", "1.49", "The DockerAPI version configured for this Docker client program")
	if err := checkCmd.MarkFlagRequired("address"); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(checkCmd)
}

func Check(cmd *cobra.Command, args []string) {
	dockerClient, err := utils.CreateDockerClient(address, clientversion, proxyaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	info, err := dockerClient.Info(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`Connection test successful!`)
	log.Printf("%v, %v, %v, %v, %v", info.OperatingSystem, info.Name, info.ServerVersion, info.OSType, info.DockerRootDir)
}
