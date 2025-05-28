package cmd

import (
	"context"
	"docker-api/utils"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

var execctrCmd = &cobra.Command{
	Use:   "execctr",
	Short: "Exec command in an existing container",
	Run:   ExecCtr,
}

func init() {
	execctrCmd.Flags().StringVar(&address, "address", "", "Target API opened IP and port: tcp://127.0.0.1:2375")
	execctrCmd.Flags().StringVar(&clientversion, "clientversion", "1.49", "The DockerAPI version configured for this Docker cli-prog")
	execctrCmd.Flags().StringVar(&command, "command", "id", "Command to execute")
	execctrCmd.Flags().StringVar(&containerid, "containerid", "", "Provide the container ID or name")

	if err := execctrCmd.MarkFlagRequired("address"); err != nil {
		log.Fatal(err)
	}
	if err := execctrCmd.MarkFlagRequired("containerid"); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(execctrCmd)
}

func ExecCtr(cmd *cobra.Command, args []string) {
	dockerClient, err := utils.CreateDockerClient(address, clientversion, proxyaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer dockerClient.Close()

	cmdParts := strings.Fields(command)
	fmt.Println(cmdParts)

	execConfig := container.ExecOptions{
		AttachStdout: true,
		AttachStderr: true,
		AttachStdin:  false,
		Tty:          false,
		Cmd:          cmdParts,
	}

	resp, err := dockerClient.ContainerExecCreate(context.Background(), containerid, execConfig)
	if err != nil {
		log.Fatalf("Failed to create exec instance: %v", err)
	}
	execID := resp.ID

	log.Printf("Exec instance created successfully, ID: %s", execID)

	execAttach, err := dockerClient.ContainerExecAttach(context.Background(), execID, container.ExecAttachOptions{})
	if err != nil {
		log.Fatal(err)
	}
	defer execAttach.Close()

	fmt.Println(`———————— Start reading exec command output ————————`)

	_, err = io.Copy(os.Stdout, execAttach.Reader)
	if err != nil && err != io.EOF {
		log.Fatalf("Error occurred while reading exec output: %v", err)
	}
	fmt.Println(`———————— End of exec command output ————————`)

	inspectResponse, err := dockerClient.ContainerExecInspect(context.Background(), execID)
	if err != nil {
		log.Fatalf("Failed to inspect exec instance status: %v", err)
	}

	log.Printf("Exec command exit code: %d", inspectResponse.ExitCode)

	if inspectResponse.ExitCode == 0 {
		log.Println("Exec command executed successfully.")
	} else {
		log.Printf("Exec command failed with non-zero exit code (%d).", inspectResponse.ExitCode)
	}
}
