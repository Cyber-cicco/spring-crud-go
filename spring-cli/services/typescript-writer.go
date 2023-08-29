package services

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/javanalyser"
	"fr.cybercicco/springgo/spring-cli/templates/angular"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func WriteAngularServiceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.ControllerPackage.Suffix + ".java", createTsService)
}

func createTsService(fileContent string){
    tokens := javanalyser.LexFile(&fileContent)
    javanalyser.PrintTokens(tokens)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    javanalyser.PrintClassAttributes(javaFile)
    javanalyser.PrintClassMethods(javaFile)
}

func createTsInterface(fileContent string){
    tokens := javanalyser.LexFile(&fileContent)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    attributes := javanalyser.GetClassAttributes(javaFile)
    interfaceName := utils.RemoveSuffix(javanalyser.GetClassName(javaFile), config.CONFIG.DtoPackage.Suffix)
    paramsMap := map[string]string{
        "{%class_name%}": interfaceName,
        "{%imports%}": "",
    }
    var attributesString = ""
    fmt.Println("----", interfaceName, "----")
    for _, attribute := range attributes{
        paramsMap["{%attribute_name%}"] = utils.ToAttributeName(attribute.Name)
        paramsMap["{%attribute_type%}"] = javanalyser.FindTsType(attribute.JavaType, paramsMap, interfaceName)
        attributesString += utils.FormatString(paramsMap, angular.INTERFACE_ATTRIBUTE_TEMPLATE)
    }
    paramsMap["{%attributes%}"] = attributesString
    daos.WriteSimpleFile(config.CONFIG.TsInterfaceFolder, utils.ToInterfaceFileName(interfaceName),[]byte(utils.FormatString(paramsMap, angular.INTERFACE_TEMPLATE))) 
}

func WriteAngularInterfaceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.DtoPackage.Suffix + ".java", createTsInterface)
}
