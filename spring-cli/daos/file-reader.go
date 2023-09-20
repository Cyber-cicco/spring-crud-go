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

/*
* fileFunc is a function that takes a string as parameter
*/
type fileFunc func(fileName string)

/* ReadJavaFileBySuffix
* Read all files in the current directory and subdirectories
* and applies a function to the ones with the given suffix
* @param suffix: the suffix of the files to read
* @param executable: the function to apply to the files
*/
func ReadJavaFileBySuffix(suffix string, executable fileFunc) {
    path := config.CONFIG.BaseJavaDir
	files, err := ioutil.ReadDir(path)
    utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
    parseFolders(files, path, suffix, executable)
}

/* FindBasePackage
*  Find the base package of the java project
*  @return the base package as a string
*/
func FindBasePackage() string {
    path := "./src/main/java"
    files, err := ioutil.ReadDir(path)
    fmt.Printf("files: %v\n", files)
    utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
    packageAsDir := ""
    findBasePackage(path + "/", files, &packageAsDir)
    packageAsDir, _ = strings.CutSuffix(packageAsDir, "/")
    packageAsDir = strings.ReplaceAll(packageAsDir, "/", ".")
    return packageAsDir

}

/* findBasePackage
*  Recursively searches for Java files within a directory and updates the base package path.
*
* This function recursively explores the directory structure starting from 'directory'.
* It identifies Java files and extracts the base package path based on the directory structure.
* The base package path is updated in the 'packageAsDir' string.
* Note that 'packageAsDir' should be initialized before calling this function.
*
* Important Note:
*   This function assumes that Java files have a '.java' extension.
*   The 'packageAsDir' string will be updated if a Java file is found.
*   If no Java files are found, 'packageAsDir' will remain unchanged.

*  @param : directory (string): The current directory path to search.
*  @param : fileInfos ([]fs.FileInfo): A list of FileInfo objects representing files and directories within 'directory'.
*  @param : packageAsDir (*string): A pointer to a string that will be updated with the base package path.
*/
func findBasePackage(directory string, fileInfos []fs.FileInfo, packageAsDir *string) {
    for _, file := range fileInfos {
        if strings.HasSuffix(file.Name(), ".java") {
            *packageAsDir = strings.ReplaceAll(directory, "./src/main/java/", "")
        }
    }
    for _, file := range fileInfos {
        if file.IsDir() {
            files, err := ioutil.ReadDir(directory + file.Name())
            utils.HandleTechnicalError(err, config.ERR_CURRENT_DIR_OPEN)
            findBasePackage(directory  + file.Name() + "/", files, packageAsDir)
        }
    }
}

/* hasSuffix
*  @return true if the file has the given suffix, false otherwise
*/
func hasSuffix(file fs.FileInfo, suffix string) bool{
    return strings.HasSuffix(file.Name(), suffix)
}

/* parseFolders 
* recursively processes files and directories within a specified path.
* 
* 
* This function recursively explores the directory structure starting from the provided 'path'.
* For each file encountered, it checks whether it is a directory or a file with the specified suffix.
* If the file is a directory, it continues to recursively process its contents.
* If the file matches the specified suffix, it reads its content and applies the 'executable' function to it.
* 
* Parameters 'path' and 'suffix' determine the scope of the search, and 'executable' defines the action to be performed on matching files.
* 
* @param files ([]fs.FileInfo): A list of FileInfo objects representing files and directories.
* @param path (string): The current directory path being processed.
* @param suffix (string): A file extension suffix used to filter files (e.g., ".txt", ".java").
* @param executable (fileFunc): A function that accepts a string as input and performs a specific operation.
*/
func parseFolders(files []fs.FileInfo, path, suffix string, executable fileFunc){
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

/* ReadTemplateField 
*  reads the content of a specified file in the "templates/java/" directory and returns it as a string.
*
*  This function takes a 'fileName' parameter, which is a string representing the name of the file to be read.
*  The function then trims any trailing newline character from the content before returning it.
*
*  @param fileName (string): The name of the file to be read, located in the "templates/java/" directory.
*  @return (string): The content of the file as a string, with any trailing newline characters removed.
*/func ReadTemplateField(fileName string) string {
    content, err := os.ReadFile("./templates/java/"+fileName)
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    return strings.TrimSuffix(string(content), "\n")
}
