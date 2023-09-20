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
    paramsMap["{%imports%}"] = ""
    paramsMap["{%urls%}"] = createServiceUrls(&serviceStruct)
    paramsMap["{%class_name%}"] = serviceStruct.Name
    paramsMap["{%http%}"] = createHttpBody(&serviceStruct, paramsMap)
    fileContent = utils.FormatString(paramsMap, angular.SERVICE_TEMPLATE)
    daos.WriteSimpleFile(config.CONFIG.TsServiceFolder, utils.ToInterfaceFileName(serviceStruct.Name),[]byte(fileContent)) 
}

func createHttpBody(serviceStruct *AngularService, paramsMap map[string]string) string {
    methods := make([]string, len(serviceStruct.Http))
    for i, method := range serviceStruct.Http {
        paramsMap["{%request_params%}"] = ""
        paramsMap["{%body%}"] = ""
        paramsMap["{%url_changer%}"] = "" 
        paramsMap["{%return_type%}"] = method.ReturnType 
        paramsMap["{%url_changed%}"] = "this." + method.Url.VarName
        paramsMap["{%request_params%}"] = "" 
        paramsMap["{%body%}"] = ""
        paramsMap["{%method%}"] = method.HttpVerb 
        paramsMap["{%method_details%}"] = utils.CreateMethodNameFromUrl(method.Url.Path)
        paramsMap["{%required_args%}"] = createServiceArgs(method, paramsMap)
        methods[i] = utils.FormatString(paramsMap, angular.SERVICE_METHOD_TEMPLATE)
    }
    return strings.Join(methods, "\n")
}

func createServiceArgs(method HttpMethod, paramsMap map[string]string) string {
    args := make([]string, len(method.Args))
    for i, arg := range method.Args {
        paramsMap["{%type%}"] = arg.Type
        paramsMap["{%name%}"] = arg.Name
        switch arg.Scope {
            case "PathVariable" : {
                paramsMap["{%url_changer%}"] = createUrlChanger(arg, method)
                paramsMap["{%url_changed%}"] = "newURL" 
            }
            case "RequestBody" : {
                paramsMap["{%body%}"] = ", " + arg.Name
                _, ok := paramsMap[arg.Type]
                if !ok{
                    paramsMap[arg.Type] = arg.Type
                    importMap := map[string]string{
                        "{%new_import%}" : arg.Type,
                        "{%file_import%}" : utils.RemoveSuffix(utils.ToInterfaceFileName(arg.Type), ".ts"),
                    }
                    paramsMap["{%imports%}"] +=  utils.FormatString(importMap, angular.SERVICE_IMPORT_TEMPLATE)
                }
            }
            case "RequestParam" : {
                paramsMap["{%request_params%}"] = "+'?" + arg.Name + "=' + " + arg.Name
            }
        }
        args[i] = utils.FormatString(paramsMap, angular.PARAMETER_TEMPLATE)
    }
    return strings.Join(args, ", ")
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
            "{%path%}" : url.AbsolutePath,
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
    fmt.Printf("config.CONFIG.TsInterfaceFolder: %v\n", config.CONFIG.TsInterfaceFolder)
    daos.WriteSimpleFile(config.CONFIG.TsInterfaceFolder, utils.ToInterfaceFileName(interfaceName),[]byte(utils.FormatString(paramsMap, angular.INTERFACE_TEMPLATE))) 
}

func WriteAngularInterfaceFile(){
    daos.ReadJavaFileBySuffix(config.CONFIG.DtoPackage.Suffix + ".java", createTsInterface)
}
