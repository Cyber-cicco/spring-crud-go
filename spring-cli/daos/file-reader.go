package daos

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/utils"
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
            fmt.Printf("file.Name(): %v\n", file.Name())
            executable(string(content))
        }
    }
}


