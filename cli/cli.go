package cli

import (
	"fmt"

	"github.com/VitJRBOG/test_vk_user_check/datamanager"
)

// ShowCLI запрашивает у пользователя ввод данных через терминал,
// передает эти данные в обработчик
// и выводит в терминал полученную информацию
func ShowCLI() {
	outputUserData(datamanager.GetUserData(inputURLToUserPage()))
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
