/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"rdotgo/cmd/utils"

	"github.com/spf13/cobra"
)

func setupZsh() {
	utils.RunSilentOsCommand("readlink /proc/$$/exe | sed \"s/.*\\///\"")
	// if currentShell != "zsh" {
	// 	panic("Please switch to zsh shell to continue")
	// }

	utils.RunOsCommand("curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh > ~/.oh-my-installer && chmod +x ~/.oh-my-installer && ~/.oh-my-installer -y")
	utils.GitClone("https://github.com/zsh-users/zsh-autosuggestions.git", "~/.oh-my-zsh/plugins/zsh-autosuggestions")
}

func setupNode() {
	utils.RunOsCommand("npm i -g n")
	utils.RunOsCommand("n 20")
}

func installNpmPackages() {
	utils.InstallGlobalNpmPackages([]string{
		"yarn",
		"eslint",
		"typescript-language-server",
		"typescript",
		"ts-node",
	})
}

func setupNvim() {
	absoluteHomeDirPath := utils.GetHomeAbsoluteDirPath()
	utils.RunOsCommand("rm -rf ~/neovim")

	utils.RunOsCommand(fmt.Sprintf("mkdir -p %s/.local/share/nvim", absoluteHomeDirPath))
	lazyPath := fmt.Sprintf("%s/.local/share/nvim/lazy/lazy.nvim", absoluteHomeDirPath)

	utils.RunOsCommand(fmt.Sprintf("git clone --filter=blob:none https://github.com/folke/lazy.nvim.git --branch=stable %s", lazyPath))

	utils.GitClone("https://github.com/neovim/neovim.git", fmt.Sprintf("%s/neovim", absoluteHomeDirPath))

	utils.RunOsCommand(`cd ~/neovim && make -j 20`)
	utils.RunOsCommand(`cd ~/neovim && sudo make install`)
}

func setupDotfiles() {
	absoluteHomeDirPath := utils.GetHomeAbsoluteDirPath()

	excludedSymlinkDirs := []string{"bin", ".tmux.conf"}
	utils.GitClone("https://github.com/circles-00/dotfiles.git", "~/dotfiles")

	files, err := os.ReadDir(fmt.Sprintf("%s/dotfiles", absoluteHomeDirPath))
	if err != nil {
		log.Fatal(err)
	}

	utils.RunOsCommand("mkdir -p ~/.config")

	fmt.Printf("FILES FILES %d", len(files))
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".git") {
			continue
		}

		if file.IsDir() && !slices.Contains(excludedSymlinkDirs, file.Name()) {
			utils.RunOsCommand(fmt.Sprintf("ln -s %s/dotfiles/%s %s/.config", absoluteHomeDirPath, file.Name(), absoluteHomeDirPath))
		}
	}

	utils.RunOsCommand(fmt.Sprintf("ln -s %s/dotfiles/tmux %s/.tmux", absoluteHomeDirPath, absoluteHomeDirPath))
	utils.RunOsCommand(fmt.Sprintf("ln -s %s/dotfiles/tmux/.tmux.conf %s/.tmux.conf", absoluteHomeDirPath, absoluteHomeDirPath))

	scripts, err := os.ReadDir(fmt.Sprintf("%s/dotfiles/bin", absoluteHomeDirPath))
	if err != nil {
		log.Fatal(err)
	}

	utils.RunOsCommand("mkdir -p ~/.local/bin")

	for _, script := range scripts {
		utils.RunOsCommand(fmt.Sprintf("ln -s %s/dotfiles/bin/.local/scripts/%s %s/.local/bin", absoluteHomeDirPath, script.Name(), absoluteHomeDirPath))
	}
}

func RunCore() {
	setupZsh()
	setupNode()
	installNpmPackages()
	setupNvim()
	setupDotfiles()
}

// coreCmd represents the core command
var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Setup all core tools",
	Long:  `Setup zsh, node, core npm packages, nvim & dotfiles`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCore()
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
