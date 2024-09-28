package cmd

import (
	"fmt"
	"strings"

	"rdotgo/cmd/utils"

	"github.com/spf13/cobra"
)

func runUbuntuUpdate() {
	utils.RunOsCommand("apt-get update -y")
	utils.RunOsCommand("apt-get upgrade -y")
	utils.RunOsCommand("apt-get install sudo -y")
	utils.RunOsCommandAsSudo("apt-get install software-properties-common -y")
}

func installCorePackages() {
	packages := []string{
		"build-essential",
		"cmake",
		"pkg-config",
		"libpthread-stubs0-dev",
		"git",
		"lua5.1",
		"unzip",
		"libtool",
		"libtool-bin",
		"gettext",
		"compton",
		"curl",
		"htop",
		"golang",
		"lsof",
		"i3",
		"ccache",
		"ninja-build",
		"python3-pip",
		"dconf-editor",
		"pavucontrol",
		"moreutils",
		"clangd",
		"ubuntu-mate-desktop",
		"flameshot",
		"tmux",
		"wireshark",
		"fzf",
		"kdenlive",
		"gimp",
		"xclip",
		"screenkey",
		"tldr",
		"ripgrep",
		"shutter",
		"software-properties-common",
		"apt-transport-https",
		"zsh",
		"luarocks",
		"nodejs",
		"npm",
	}

	utils.RunOsCommandAsSudo(
		fmt.Sprintf("DEBIAN_FRONTEND=noninteractive apt-get install -y %s", strings.Join(packages, " ")),
	)
}

func setupChromeBrowser() {
	utils.RunOsCommand("cd ~ && wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb")
	utils.RunOsCommandAsSudo("apt-get install -y ~/google-chrome-stable_current_amd64.deb")
	utils.RunOsCommandAsSudo("apt --fix-broken install -y")
}

func RunUbuntuCmd() {
	runUbuntuUpdate()
	installCorePackages()
	setupChromeBrowser()
}

// dependenciesCmd represents the ubuntu command
var dependenciesCmd = &cobra.Command{
	Use:   "ubuntu",
	Short: "The ubuntu command will install the needed dependencies for your Ubuntu machine",
	Long:  `The ubuntu command will install the needed dependencies for your Ubuntu machine`,
	Run: func(cmd *cobra.Command, args []string) {
		RunUbuntuCmd()
	},
}

func init() {
	rootCmd.AddCommand(dependenciesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ubuntuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ubuntuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
