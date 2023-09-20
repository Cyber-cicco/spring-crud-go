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

/* writeFile
*  writeFile checks if a file with the specified 'filename' already exists. If it does, it overrides the file with the provided 'bytes'.
*  If the file does not exist, it creates a new file with the specified 'filename' and writes the provided 'bytes' to it.
*
*  @param bytes ([]byte): The content to be written to the file.
*  @param filename (string): The name of the file to be written.
*/
func writeFile(bytes []byte, filename string) {
	if FileExists(filename) {
		overrideFile(bytes, filename)
	} else {
		createNewFile(bytes, filename)
	}
}

/* createNewFile
* creates a new file with the specified 'filename' and writes the provided 'bytes' to it.
*
* @param bytes ([]byte): The content to be written to the file.
* @param filename (string): The name of the file to be created.
*/func createNewFile(bytes []byte, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
    defer f.Close()
	utils.HandleTechnicalError(err, config.ERR_FILE_CREATION)
	f.Write(bytes)
	f.Sync()
}

/* overrideFile
*  overrides an existing file with the provided 'bytes', if the configuration allows it.
*
*  If the 'EreaseFiles' configuration is set to false, the function prints a warning message and does not override the file.
*  Otherwise, it truncates the file, writes the new content, and closes the file.
*
*  @param bytes ([]byte): The content to be written to the file.
*  @param filename (string): The name of the file to be overridden.
*/
func overrideFile(bytes []byte, filename string) {
	if !config.CONFIG.EreaseFiles {
		fmt.Println("A file was at risk of being overwritten. If this was the intention, change the settings in spring-parameters.json")
	} else {
		f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.Write(bytes)
	}
}
/* FileExists 
*  checks if a file or directory exists at the specified 'dirPath'.
*  It returns 'true' if the file or directory exists, and 'false' otherwise.
*
*  @param dirPath (string): The path of the file or directory to be checked for existence.
*  @return (bool): 'true' if the file or directory exists, 'false' otherwise.
*/
func FileExists(dirPath string) bool {
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

/* WriteSimpleFile 
*  writes the provided 'content' to a file with the specified 'filename' in the specified 'directory'.
*
*  If the specified 'directory' does not exist, it attempts to create it with read, write, and execute permissions for all users.
*  It then calls the 'writeFile' function to perform the file writing operation.
*
*  @param directory (string): The directory where the file will be created.
*  @param filename (string): The name of the file to be created.
*  @param content ([]byte): The content to be written to the file.
*/func WriteSimpleFile(directory, filename string, content []byte) {
	if !FileExists(directory) {
		utils.HandleTechnicalError(os.MkdirAll(directory, 0777), config.ERR_DIR_CREATION)
	}
	writeFile(content, directory+filename)
}

/* WriteTemplateFile 
* generates and writes various components of a Java class represented by the provided 'javaClass' object.
*
* This function takes two parameters:
*   - 'javaClass': A BaseJavaClass object representing a Java class with different components (e.g., Imports, Annotations, etc.).
*   - 'directory': A string specifying the directory where the template files will be written.
*
* It constructs the directory path in the "templates/java/" directory and uses 'WriteSimpleFile' to write each component to a corresponding file.
*
* @param javaClass (entities.BaseJavaClass): The Java class containing different components (Imports, Annotations, etc.).
* @param directory (string): The directory within "templates/java/" where the template files will be written.
*/
func WriteTemplateFile(javaClass entities.BaseJavaClass, directory string){
    directory = "templates/java/" + directory
    WriteSimpleFile(directory,"Imports.txt", []byte(javaClass.Imports ))
    WriteSimpleFile(directory,"Annotations.txt", []byte(javaClass.Annotations))
    WriteSimpleFile(directory,"ClassType.txt", []byte(javaClass.ClassType))
    WriteSimpleFile(directory,"Extends.txt", []byte(javaClass.Extends))
    WriteSimpleFile(directory,"Implements.txt", []byte(javaClass.Implements))
    WriteSimpleFile(directory,"Body.txt", []byte(javaClass.Body))
}

/* WriteBaseConfigFile 
*  writes the provided base configuration 'baseConfig' to the "spring-parameters.json" file.
*
*  It serializes the configuration into JSON format, and writes the JSON content to the file.
*  If the directory "./src/main/java/" does not exist, it raises a usage error.
*
*  @param baseConfig (config.Config): The base configuration information.
*/
func WriteBaseConfigFile(baseConfig config.Config) {
	confBytes, err := json.MarshalIndent(baseConfig, "", "    ")
	utils.HandleTechnicalError(err, config.ERR_MARSHARL)
    if !FileExists("./src/main/java/"){
        utils.HandleUsageError(errors.New("io error"), config.ERR_NO_JAVA)
    }
	writeFile(confBytes, "./spring-parameters.json")
}

/* DeleteJpaFiles 
*  deletes all the entity configuration files in the JPA directory specified in the configuration.
*
*  It reads the files in the JPA directory, then iterates through them, removing each one.
*
*  Note: This function does not handle errors related to file deletion.
*/
func DeleteJpaFiles() {
	files, err := ioutil.ReadDir(config.CONFIG.JpaJsonFilePath)
    utils.HandleTechnicalError(err, config.ERR_JPA_DIR_OPEN)
    for _, file := range files {
        os.Remove(config.CONFIG.JpaJsonFilePath + file.Name()) 
    }
}
