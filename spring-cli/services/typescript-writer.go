package services

import (
	"fmt"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/javanalyser"
	"fr.cybercicco/springgo/spring-cli/templates/angular"
	"fr.cybercicco/springgo/spring-cli/utils"
)


func WriteAngularServiceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.ControllerPackage.Suffix + ".java", createTsService2)
}

func createTsService2(fileContent string){
    tokens := javanalyser.LexFile(&fileContent)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    mapJavaService(javaFile)
}

func createTsService(fileContent string){
    tokens := javanalyser.LexFile(&fileContent)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    classPath := javanalyser.GetClassPath(javaFile)
    for _, javaMethod := range javaFile.JavaClass.Methods{
        httpVerb := javanalyser.FindHttpVerb(javaMethod) 
        if httpVerb != ""{
            createOneMethod(&classPath, &httpVerb, javaFile, javaMethod)
        }
    }
}

func createOneMethod(classPath, httpVerb *string, javaFile javanalyser.JavaInterpreted, method javanalyser.Method){
        paramsMap := map[string]string{}
        methodPath := javanalyser.GetMethodPath(method, *classPath)
        prepareUrlMap(javaFile, method, &methodPath)
        prepareHttpMap(javaFile, method, *httpVerb)
        fmt.Println(methodPath)
        fmt.Println(paramsMap)
}

func prepareUrlMap(javaFile javanalyser.JavaInterpreted, method javanalyser.Method, methodPath *string){
}


func prepareHttpMap(javaFile javanalyser.JavaInterpreted, method javanalyser.Method, httpVerb string) {
    paramsMap := map[string]string{}
    paramsMap["{%target_name%}"] = javanalyser.FindTsType(method.ReturnType, paramsMap, javaFile.JavaClass.Name.Name.Value)
    paramsMap["{%method%}"] = httpVerb
    paramsMap["{%by%}"] = ""
    paramsMap["{%target_name%}"]  = createParameters(method, paramsMap)

}

func createParameters(method javanalyser.Method, paramsMap map[string]string) string {
    parameters := []string{}
    paramsMap["{%url_changer%}"] = ""
    for _, variable := range method.Parameters{
        createUrlChanger(variable, paramsMap)    
        paramsMap["{%name%}"] = variable.Name.Value
        paramsMap["{%type%}"] = javanalyser.FindTsType(variable.JavaType, paramsMap, "")
        parameters = append(parameters, utils.FormatString(paramsMap, angular.PARAMETER_TEMPLATE))
    }
    return strings.Join(parameters, ", ")
}

func createUrlChanger(variable javanalyser.Variable, paramsMap map[string]string){
    for _, annotation := range variable.Annotations {
        if annotation.Name.Name.Value == "PathVariable"{
            
        }
    }
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
