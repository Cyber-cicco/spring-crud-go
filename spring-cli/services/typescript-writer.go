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
    daos.ReadJavaFileBySuffix(config.CONFIG.ControllerPackage.Suffix + ".java", createTsService)
}

func createTsService(fileContent string){
    paramsMap := map[string]string{}
    tokens := javanalyser.LexFile(&fileContent)
    javaFile := javanalyser.OrganizeTokensByMeaning(tokens)
    serviceStruct := mapJavaService(javaFile)
    paramsMap["{%urls%}"] = createServiceUrls(&serviceStruct)
    paramsMap["{%class_name%}"] = serviceStruct.Name
    paramsMap["{%http%}"] = createHttpBody(&serviceStruct, paramsMap)
    fmt.Printf("paramsMap: %v\n", paramsMap)
}

func createHttpBody(serviceStruct *AngularService, paramsMap map[string]string) string {
    methods := make([]string, len(serviceStruct.Http))
    for i, method := range serviceStruct.Http {
        paramsMap["{%by%}"] = ""
        paramsMap["{%request_params%}"] = ""
        paramsMap["{%body%}"] = ""
        paramsMap["{%url_changer%}"] = ""
        paramsMap["{%return_type%}"] = ""
        paramsMap["{%url_changed%}"] = ""
        paramsMap["{%request_params%}"] = ""
        paramsMap["{%body%}"] = ""
        paramsMap["{%method%}"] = method.HttpVerb
        paramsMap["{%required_args%}"] = createServiceArgs(method, paramsMap)
        methods[i] = utils.FormatString(paramsMap, angular.SERVICE_METHOD_TEMPLATE)
    }
    return ""
}

func createServiceArgs(method HttpMethod, paramsMap map[string]string) string {
    args := make([]string, len(method.Args))
    for i, arg := range method.Args {
        paramsMap["{%type%}"] = arg.Type
        paramsMap["{%name%}"] = arg.Name
        if arg.Scope == "PathVariable" {
            paramsMap["{%url_changer%}"] = createUrlChanger(arg, method)
        }
        args[i] = utils.FormatString(paramsMap, angular.PARAMETER_TEMPLATE)
    }
    return ""
}

func createUrlChanger(arg Arg, method HttpMethod) string {
    paramsMap := map[string]string {
        "{%url%}" : method.Url.VarName,
        "{%match%}" : arg.Name,
    }
    return utils.FormatString(paramsMap, angular.URL_CHANGER)
}

func createServiceUrls(serviceStruct *AngularService) string {
    urls := make([]string, len(serviceStruct.Urls))
    for i, url := range serviceStruct.Urls {
        paramsMap := map[string]string{
            "{%url_var%}" : url.VarName,
            "{%path%}" : url.Path,
        }
        urls[i] = utils.FormatString(paramsMap, angular.URL_DECLARATION)
    }
    return strings.Join(urls, "\n")

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
