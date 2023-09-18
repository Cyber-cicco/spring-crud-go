package daos

import (
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/utils"
	"fr.cybercicco/springgo/spring-cli/entities"
)

type angularFunc func(fileName string)

func ReadJavaFileBySuffix(suffix string, executable angularFunc) {
    path := config.CONFIG.BaseJavaDir
	files, err := ioutil.ReadDir(path)
    utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
    parseFolders(files, path, suffix, executable)
}

func hasSuffix(file fs.FileInfo, suffix string) bool{
    return strings.HasSuffix(file.Name(), suffix)
}

func parseFolders(files []fs.FileInfo, path, suffix string, executable angularFunc){
    for _, file := range files {
        if file.IsDir() {
            files, err := ioutil.ReadDir(path+"/"+file.Name())
            utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
            parseFolders(files, path+"/"+file.Name(), suffix, executable)
        } else if hasSuffix(file, suffix) {
            content, err := os.ReadFile(path+"/"+file.Name())
            utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
            executable(string(content))
        }
    }
}

func ReadTemplateFile([]entities.BaseJavaClass){
    fileExists("template/")
}
func ReadTemplateField(fileName string) string {
    content, err := os.ReadFile("./templates/java/"+fileName)
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    return strings.TrimSuffix(string(content), "\n")
}
