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

func LoadEntityJson() ([]entities.JpaEntity, error){
    directoryPath := config.CONFIG.JpaJsonFilePath 
    files, err := ioutil.ReadDir(directoryPath)
    utils.HandleBasicError(err, "Erreur dans l'ouverture du dossier censé contenir les fichiers de configuration des entités JPA")
    noEntitiesErr error = nil
    jpaEntities := []entities.JpaEntity{}
    for _, file := range files {
        fmt.Println(file.Name())
        var jpaEntity entities.JpaEntity
        data, fileErr := ioutil.ReadFile(directoryPath + file.Name())
        utils.HandleBasicError(fileErr, "Erreur dans la lecture de fichier de configuration d'une entité JPA")
        utils.HandleBasicError(json.Unmarshal(data, &jpaEntity), "Erreur dans la lecture des ")
        jpaEntities = append(jpaEntities, jpaEntity)
    }
    return jpaEntities[:], noEntitesErr
}
