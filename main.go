package main

import (
	"github.com/VitJRBOG/VkUserChecker/cli"
	"github.com/VitJRBOG/VkUserChecker/filemanager"
)

func main() {
	cli.ShowCLI(filemanager.GetConfig())
}
