package utils

import "fmt"

func SetupUbuntuMachine() {
	containerName := "rdotgo-ubuntu-container"
	// RunOsCommand(fmt.Sprintf("docker container stop %s && docker container rm %s && docker run -d --name rdotgo-ubuntu-container ubuntu:latest tail -f /dev/null", containerName, containerName))
	RunOsCommand(fmt.Sprintf("docker cp ~/go/bin/rdotgo %s:/root/rdotgo", containerName))

	RunOsCommandWithStdIn(fmt.Sprintf("docker exec -it %s /bin/bash", containerName))
}
