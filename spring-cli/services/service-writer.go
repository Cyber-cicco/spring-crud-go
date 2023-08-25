package services

import (

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/javanalyser"
)

func WriteAngularServiceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.ControllerPackage.Suffix + ".java", createTsService)
}

func createTsService(fileContent string){
    tokens := javanalyser.LexFile(&fileContent)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    javanalyser.PrintFile(javaFile)
}

