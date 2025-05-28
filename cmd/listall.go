package cmd

import (
	"context"
	"docker-api/utils"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/spf13/cobra"
	"log"
)

var listallCmd = &cobra.Command{
	Use:   "listall",
	Short: "List all containers and images",
	Run:   ListAll,
}

func init() {
	listallCmd.Flags().StringVar(&address, "address", "", "Target API opened IP and port: tcp://127.0.0.1:2375")
	listallCmd.Flags().StringVar(&clientversion, "clientversion", "1.49", "The DockerAPI version configured for this Docker client program")
	if err := listallCmd.MarkFlagRequired("address"); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(listallCmd)
}

func ListAll(cmd *cobra.Command, args []string) {
	dockerClient, err := utils.CreateDockerClient(address, clientversion, proxyaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	fmt.Println(`———————— ContainerList ————————`)
	containerList, err := dockerClient.ContainerList(context.Background(), container.ListOptions{All: true})

	if err != nil {
		log.Fatal(err)
	}
	for _, ctr := range containerList {
		fmt.Println(ctr.ID[:12], ctr.Names, ctr.Image, ctr.Status)
	}

	fmt.Println(`———————— ImageList ————————`)

	imageList, err := dockerClient.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range imageList {
		fmt.Println(image.ID[7:19], image.RepoTags, image.Size, image.Containers)
	}

}
