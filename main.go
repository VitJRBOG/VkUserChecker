package main

import (
	"github.com/VitJRBOG/test_vk_user_check/cli"
	"github.com/VitJRBOG/test_vk_user_check/filemanager"
)

func main() {
	cli.ShowCLI(filemanager.GetConfig())
}

// func test() {
// 	cfg, _ := filemanager.GetConfig()
// 	fmt.Println(datamanager.GetUserData(cfg, "https://vk.com/id1"))
// }
