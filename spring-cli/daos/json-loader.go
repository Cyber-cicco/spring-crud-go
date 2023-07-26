package daos

import (
	"encoding/json"
	"fmt"
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
    utils.HandleBasicError(fileErr, "Erreur dans l'ouverture du fichier de configuration")
    utils.HandleBasicError(json.Unmarshal(data, &config.CONFIG), "Erreur dans la lecture du fichier de configuration")
    config.CONFIG.BaseJavaDir = config.RELATIVE_PATH + config.JAVA_PATH + strings.ReplaceAll(config.CONFIG.BasePackage, ".", "/") + "/"
    config.CONFIG.JpaJsonFilePath =  "../jpa/"
    utils.HandleBasicError(fileExists(config.CONFIG.BaseJavaDir), "Erreur : le package précisé dans le fichier de configuration ne semble pas pointer vers un répertoire existant.")
    fmt.Println(config.CONFIG)
}


func Jsonify(entity entities.JpaEntity) []byte{
    entityJson, err := json.MarshalIndent(entity, "", "    ")
    unformattedJson := string(entityJson)
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003c", "<")
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003e", ">")
    utils.HandleBasicError(err, "Erreur interne : ")
    return []byte(unformattedJson)
}


