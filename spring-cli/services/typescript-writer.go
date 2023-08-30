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
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    classPath := javanalyser.GetClassPath(javaFile)
    httpMethods := ""
    fmt.Println("caca")
    for _, javaMethod := range javaFile.JavaClass.Methods{
        httpVerb := javanalyser.FindHttpVerb(javaMethod) 
        if httpVerb != ""{
            createOneMethod(&classPath, &httpMethods, javaFile, javaMethod)
        }
    }
}

func createOneMethod(classPath, httpMethods *string, javaFile javanalyser.JavaInterpreted, method javanalyser.Method){
        paramsMap := prepareHttpMap(javaFile, method, *httpMethods)
        methodPath := javanalyser.GetMethodPath(javaFile, method, *classPath)
        fmt.Println(methodPath)
        fmt.Println(paramsMap)
}


func prepareHttpMap(javaFile javanalyser.JavaInterpreted, method javanalyser.Method, httpMethod string) map[string]string {
    paramsMap := map[string]string{}
    paramsMap["{%target_name%}"] = javanalyser.FindTsType(method.ReturnType, paramsMap, javaFile.JavaClass.Name.Name.Value)
    paramsMap["{%method%}"] = httpMethod
    paramsMap["{%by%}"] = ""
    for _, variable := range method.Parameters{
        fmt.Printf("variable: %v\n", variable)
    }
    return paramsMap

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
