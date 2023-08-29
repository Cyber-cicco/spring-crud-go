package javanalyser

import (
    "fmt"
    "strings"

    "fr.cybercicco/springgo/spring-cli/config"
    "fr.cybercicco/springgo/spring-cli/templates/angular"
    "fr.cybercicco/springgo/spring-cli/utils"
)

var typeMapping = map[string]string{
    "String": "string",
    "int": "number",
    "Integer": "number",
    "long": "number",
    "Long": "number",
    "float": "number",
    "double": "number",
    "Double": "number",
    "Float": "number",
    "boolean": "boolean",
    "Boolean": "boolean",
    "LocalDateTime" : "Date",
    "LocalDate" : "Date",
    "LocalTime" : "Date",
    "MultipartFile" : "File",
    "HttpStatus":"HttpStatusCode",
    // add more type mappings as needed
}

/**La méthode fonctionne pour les fields de base, les listes et les sets. Si il y a des types plus complexes, cela risque de poser problème*/
func FindTsType(javaType JavaType, paramsMap map[string]string, className string) string {
    tsType, ok := typeMapping[javaType.Name.Value]
    if ok {
        return tsType
    } else if strings.HasSuffix(javaType.Name.Value, config.CONFIG.DtoPackage.Suffix) {
        tsTypeName := javaType.Name.Value[:len(javaType.Name.Value)-len(config.CONFIG.DtoPackage.Suffix)]
        fmt.Println(tsTypeName)
        if tsTypeName != className {
            paramsMap["{%file_import%}"] = utils.ToInterfaceFileName(tsTypeName)[0:len(utils.ToInterfaceFileName(tsTypeName))-3]
            paramsMap["{%new_import%}"] = tsTypeName
            paramsMap["{%imports%}"] += utils.FormatString(paramsMap, angular.INTERFACE_IMPORT_TEMPLATE)
        }
        return tsTypeName
    } else {
        switch javaType.Name.Value {
        case "List":
            return FindTsType(javaType.SubTypes[0], paramsMap, className) + "[]"
        case "Set":
            return FindTsType(javaType.SubTypes[0], paramsMap, className) + "[]"
        case "ResponseEntity":
            return FindTsType(javaType.SubTypes[0], paramsMap, className)
        default:
            return "any"
        }
    }
}
