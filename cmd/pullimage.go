package cmd

import (
	"context"
	"docker-api/utils"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/image"
	"github.com/spf13/cobra"
	"io"
	"log"
)

var pullimageCmd = &cobra.Command{
	Use:   "pullimage",
	Short: "Pull a new image",
	Run:   PullImage,
}

func init() {
	pullimageCmd.Flags().StringVar(&address, "address", "", "Target API opened IP and port: tcp://127.0.0.1:2375")
	pullimageCmd.Flags().StringVar(&clientversion, "clientversion", "1.49", "The DockerAPI version configured for this Docker client program")
	pullimageCmd.Flags().StringVar(&imagename, "imagename", "ubuntu:latest", "Image name (tag optional)")
	if err := pullimageCmd.MarkFlagRequired("address"); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(pullimageCmd)
}

func PullImage(cmd *cobra.Command, args []string) {
	dockerClient, err := utils.CreateDockerClient(address, clientversion, proxyaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	reader, err := dockerClient.ImagePull(context.Background(), imagename, image.PullOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	for {
		var message struct {
			Status         string `json:"status"`
			Progress       string `json:"progress"`
			ProgressDetail struct {
				Current int `json:"current"`
				Total   int `json:"total"`
			} `json:"progressDetail"`
		}
		if err := decoder.Decode(&message); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Failed to parse progress information: %v", err)
		}

		if message.Status != "" {
			if message.Progress != "" {
				fmt.Printf("%s: %s (%d/%d)\n", message.Status, message.Progress, message.ProgressDetail.Current, message.ProgressDetail.Total)
			} else {
				fmt.Println(message.Status)
			}
		}
	}
}
