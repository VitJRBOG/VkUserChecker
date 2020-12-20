package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config хранит параметры для программы
type Config struct {
	AccessToken string      `json:"access_token"`
	Communities []Community `json:"communities"`
}

// Community хранит информацию о сообществе
type Community struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SetConfig записывает новые параметры в файл config.json
func SetConfig(cfgValues Config) {
	writeJSONToFile(getPathToConfigFile(), makeJSON(cfgValues))
}

// GetConfig получает параметры из файла config.json,
// если файл config.json не будет обнаружен рядом с исполняемым файлом программы,
// то будет создан новый, а пользователь получит уведомление об этом
func GetConfig() (Config, bool) {
	return readConfigFile()
}

func checkConfigFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		makeConfigFile(path)
		return true
	}
	return false
}

func getPathToConfigFile() string {
	pathToCurrentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	pathToConfigFile := filepath.FromSlash(pathToCurrentDir + "/config.json")
	return pathToConfigFile
}

func makeConfigFile(path string) {
	var communitiesValues []Community
	communitiesValues = append(communitiesValues, Community{1, "ApiClub"})
	var cfgValues = Config{"", communitiesValues}

	writeJSONToFile(path, makeJSON(cfgValues))
}

func makeJSON(cfgValues Config) []byte {
	content, err := json.Marshal(cfgValues)
	if err != nil {
		panic(err.Error())
	}
	return content
}

func writeJSONToFile(path string, content []byte) {
	ioutil.WriteFile(path, content, 0644)
}

func readConfigFile() (Config, bool) {
	path := getPathToConfigFile()
	fileWasCreated := checkConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return parseJSON(content), fileWasCreated
}

func parseJSON(content []byte) Config {
	var cfgValues Config
	err := json.Unmarshal(content, &cfgValues)
	if err != nil {
		panic(err.Error())
	}
	return cfgValues
}
