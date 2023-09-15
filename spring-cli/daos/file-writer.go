package daos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func writeFile(bytes []byte, filename string) {
	if fileExists(filename) {
        fmt.Printf("filename: %v\n", filename)
		overrideFile(bytes, filename)
	} else {
		createNewFile(bytes, filename)
	}
}

func createNewFile(bytes []byte, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
	utils.HandleTechnicalError(err, config.ERR_FILE_CREATION)
	f.Write(bytes)
	f.Sync()
	f.Close()
}

func overrideFile(bytes []byte, filename string) {
	if !config.CONFIG.EreaseFiles {
		fmt.Println("Un fichier risquait de se faire réécrire. Si cela était le but, changer les paramètres de votre configuration")
	} else {
		f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.Write(bytes)
	}
}

func fileExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func WriteEntityJson(entity entities.JpaEntity) {
	writeFile(entity.FileBytes, entity.FileName)
}

func WriteSimpleFile(directory, filename string, content []byte) {
	if !fileExists(directory) {
		utils.HandleTechnicalError(os.MkdirAll(directory, 0777), config.ERR_DIR_CREATION)
	}
	writeFile(content, directory+filename)
}

func WriteBaseConfigFile(baseConfig config.Config) {
	confBytes, err := json.MarshalIndent(baseConfig, "", "    ")
	utils.HandleTechnicalError(err, config.ERR_MARSHARL)
    if !fileExists("./src/main/java/"){
        utils.HandleUsageError(errors.New("io error"), config.ERR_NO_JAVA)
    }
	writeFile(confBytes, "./spring-parameters.json")
}

/*
deletes all the entity config files in the jpa directory
*/
func DeleteJpaFiles() {
	files, err := ioutil.ReadDir(config.CONFIG.JpaJsonFilePath)
    utils.HandleTechnicalError(err, config.ERR_JPA_DIR_OPEN)
    for _, file := range files {
        os.Remove(config.CONFIG.JpaJsonFilePath + file.Name()) 
    }
}
