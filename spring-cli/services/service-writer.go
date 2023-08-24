package services

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
)

func WriteAngularServiceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.ControllerPackage.Suffix + ".java", createTsService)
}

func createTsService(fileContent string){
    tokens := LexFile(&fileContent)
    for _, tokenLine := range tokens.tokens{
        for _, token := range tokenLine{
            fmt.Print(string(token.value) + " ")
        }
        fmt.Println()
    }
}
