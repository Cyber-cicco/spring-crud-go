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

/* LoadConfig 
*  reads and initializes the configuration from the "spring-parameters.json" file.
*
*  This function performs the following steps:
*    1. Reads the content of the "spring-parameters.json" file.
*    2. Unmarshals the JSON data into the 'config.CONFIG' variable.
*    3. Sets specific fields in the configuration, such as 'BaseJavaDir' and 'JpaJsonFilePath'.
*    4. Checks if the specified base Java directory and package specified in the configuration file; raises usage errors if not.
*    5. Verifies the existence of TypeScript interface and service folders; raises usage errors if not configured.
*
*  Note: This function may raise usage errors if certain configurations are missing or incorrect.
*/
func LoadConfig() {
    file := filepath.Join("spring-parameters.json")
    data, fileErr := ioutil.ReadFile(file)
    utils.HandleUsageError(fileErr, config.ERR_OPEN_CONFIG)
    utils.HandleTechnicalError(json.Unmarshal(data, &config.CONFIG), config.ERR_UNMARSHAL)
    config.CONFIG.BaseJavaDir = config.JAVA_PATH 
    config.CONFIG.JpaJsonFilePath =  "./jpa/"
    exists := FileExists(config.CONFIG.BaseJavaDir + strings.ReplaceAll(config.CONFIG.BasePackage, ".", "/") + "/")
    if !exists {
        utils.HandleUsageError(errors.New("config error"), config.ERR_BAD_CONFIG_PACKAGE)
    }
    if config.CONFIG.TsInterfaceFolder == "" || config.CONFIG.TsServiceFolder == "" {
        utils.HandleUsageError(errors.New("config error"), config.ERR_BAD_TS_DIRECTORY)
    }
}

/* Jsonify 
*  serializes a JpaEntity into a JSON byte slice, with special handling for specific characters.
*
*  It performs the following steps:
*    1. Marshals the 'entity' into a JSON byte slice with indentation.
*    2. Performs character replacement to characters needed for java types and automatically escaped by the Unmarshal from go.
*    3. Returns the resulting JSON byte slice.
*
*  @param entity (entities.JpaEntity): The JpaEntity object to be serialized.
*  @return ([]byte): The serialized JSON data as a byte slice.
*/
func Jsonify(entity entities.JpaEntity) []byte{
    entityJson, err := json.MarshalIndent(entity, "", "    ")
    utils.HandleTechnicalError(err, "Erreur interne : ")
    unformattedJson := string(entityJson)
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003c", "<")
    unformattedJson = strings.ReplaceAll(unformattedJson, "\\u003e", ">")
    return []byte(unformattedJson)
}
/* LoadEntityJson 
*  reads JpaEntity JSON files from the specified directory and returns them as a slice of JpaEntity objects.
*
*  It performs the following steps:
*    1. Reads the content of JSON files in the 'config.CONFIG.JpaJsonFilePath' directory.
*    2. Unmarshals each JSON data into a JpaEntity object.
*    3. Appends each JpaEntity to the 'jpaEntities' slice.
*    4. Returns the slice of JpaEntity objects.
*
*  Note: This function may raise usage or technical errors depending on the encountered conditions.
*
*  @return ([]entities.JpaEntity): A slice of JpaEntity objects read from JSON files.
*/
func LoadEntityJson() []entities.JpaEntity{
    directoryPath := config.CONFIG.JpaJsonFilePath 
    files, err := ioutil.ReadDir(directoryPath)
    utils.HandleUsageError(err, config.ERR_JPA_DIR_OPEN)
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
