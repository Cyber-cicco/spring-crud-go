package daos

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)


func LoadConfig() {
    file := filepath.Join(config.RELATIVE_PATH, "spring-parameters.json")
    data, fileErr := ioutil.ReadFile(file)
    utils.HandleTechnicalError(fileErr, config.ERR_OPEN_CONFIG)
    utils.HandleTechnicalError(json.Unmarshal(data, &config.CONFIG), config.ERR_UNMARSHAL)
    config.CONFIG.BaseJavaDir = config.RELATIVE_PATH + config.JAVA_PATH 
    config.CONFIG.JpaJsonFilePath =  "../jpa/"
    utils.HandleTechnicalError(fileExists(config.CONFIG.BaseJavaDir + strings.ReplaceAll(config.CONFIG.BasePackage, ".", "/") + "/"), config.ERR_BAD_CONFIG_PACKAGE)
}


func Jsonify(entity entities.JpaEntity) []byte{
    entityJson, err := json.MarshalIndent(entity, "", "    ")
    unformattedJson := string(entityJson)
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003c", "<")
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003e", ">")
    utils.HandleTechnicalError(err, "Erreur interne : ")
    return []byte(unformattedJson)
}

func LoadEntityJson() []entities.JpaEntity{
    directoryPath := config.CONFIG.JpaJsonFilePath 
    files, err := ioutil.ReadDir(directoryPath)
    utils.HandleTechnicalError(err, config.ERR_JPA_DIR_OPEN)
    if len(files) == 0 {
        utils.HandleUsageError(errors.New("files not found error"), config.ERR_NO_JPA_FILE)
    }
    jpaEntities := []entities.JpaEntity{}
    for _, file := range files {
        var jpaEntity entities.JpaEntity
        data, fileErr := ioutil.ReadFile(directoryPath + file.Name())
        utils.HandleTechnicalError(fileErr, config.ERR_JPA_FILE_READ)
        utils.HandleTechnicalError(json.Unmarshal(data, &jpaEntity), config.ERR_UNMARSHAL)
        jpaEntities = append(jpaEntities, jpaEntity)
    }
    return jpaEntities[:]
}
