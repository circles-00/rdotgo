package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunOsCommand(cmd string) {
	fmt.Printf("Running command: %s\n", cmd)
	command := exec.Command("bash", "-c", cmd)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	command.Stdout = mw
	command.Stderr = mw

	err := command.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(stdBuffer.String())
}

func RunSilentOsCommand(cmd string) string {
	command := exec.Command("bash", "-c", cmd)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	command.Stdout = mw
	command.Stderr = mw

	err := command.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return stdBuffer.String()
}

func RunOsCommandWithStdIn(cmd string) {
	command := exec.Command("bash", "-c", cmd)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func RunOsCommandAsSudo(cmd string) {
	// NOTE: Might need to pass password here tho
	RunOsCommand(fmt.Sprintf("sudo -S %s", cmd))
}

func GitClone(repo string, destination string) {
	RunOsCommand(fmt.Sprintf("git clone %s %s", repo, destination))
}

func InstallGlobalNpmPackages(packages []string) {
	RunOsCommand(fmt.Sprintf("npm i -g %s", strings.Join(packages, " ")))
}

func GetHomeAbsoluteDirPath() string {
	absoluteHomeDirPath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return absoluteHomeDirPath
}
