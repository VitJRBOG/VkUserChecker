package main

import (
	"github.com/VitJRBOG/test_vk_user_check/cli"
	"github.com/VitJRBOG/test_vk_user_check/filemanager"
)

func main() {
	cli.ShowCLI(filemanager.GetConfig())
}
