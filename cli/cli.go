package cli

import (
	"fmt"
	"os"

	"github.com/VitJRBOG/VkUserChecker/datamanager"
	"github.com/VitJRBOG/VkUserChecker/filemanager"
)

// ShowCLI запрашивает у пользователя ввод данных через терминал,
// передает эти данные в обработчик
// и выводит в терминал полученную информацию
func ShowCLI(cfgValues filemanager.Config, cfgFileWasCreated bool) {
	if cfgFileWasCreated {
		fmt.Println("File config.json has been created and have no data now.")
		os.Exit(0)
	}
	outputUserData(datamanager.GetUserData(cfgValues, inputURLToUserPage()))
}

func inputURLToUserPage() string {
	fmt.Print("--- Enter URL to user's page and press «Enter» ---\n" +
		"> ")
	var userAnswer string
	_, err := fmt.Scan(&userAnswer)
	if err != nil {
		panic(err.Error())
	}
	return userAnswer
}

func outputUserData(userData string) {
	fmt.Println(userData)
}
