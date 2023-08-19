package daos

import (
	"encoding/json"
	"fmt"
	"os"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func writeFile(bytes []byte, filename string){
    err := fileExists(filename)
    if err == nil {
        overrideFile(bytes, filename)
    } else {
        createNewFile(bytes, filename)
    }
}

func createNewFile(bytes []byte, filename string){
    f,err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0660);    
    utils.HandleBasicError(err, "Erreur dans la création d'un fichier")
    f.Write(bytes)
    f.Sync()
    f.Close()
}

func overrideFile(bytes []byte, filename string){
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

func fileExists(dirPath string) error {
    _, err := os.Stat(dirPath)
    if err != nil {
        if os.IsNotExist(err) {
            return err
        }
        return err
    }
    return nil
}

func WriteEntityJson(entity entities.JpaEntity){
    writeFile(entity.FileBytes, entity.FileName)
}

func WriteJavaClass(directory, filename string, content []byte){
    if fileExists(directory) != nil {
        utils.HandleBasicError(os.MkdirAll(directory, 0777), "Erreur dans la création d'un répertoire") 
    }
    writeFile(content, directory + filename)
}

func WriteBaseConfigFile(baseConfig config.Config) {
    confBytes, err := json.MarshalIndent(baseConfig, "", "    ")
    utils.HandleBasicError(err, "Erreur dans la désérialisation d'un objet en JSON")
    fmt.Println("caca")
    writeFile(confBytes, "test.json")
}
