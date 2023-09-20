package javanalyser

import (
    "strings"

    "fr.cybercicco/springgo/spring-cli/config"
    "fr.cybercicco/springgo/spring-cli/templates/angular"
    "fr.cybercicco/springgo/spring-cli/utils"
)

//Map of java types to typescript types
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

/*FindTsType 
* maps a Java type to its corresponding TypeScript type.
*
* The function first checks if there is a direct mapping for the Java type in 'typeMapping'.
* If found, it returns the corresponding TypeScript type.
* If not, it checks if the Java type is a DTO (Data Transfer Object) by suffix matching.
* If it's a DTO, it updates 'paramsMap' with import information and returns the DTO type name.
* If not, it handles special cases like Lists, Sets, and ResponseEntity, and returns appropriate TypeScript types.
* If no mapping is found, it defaults to 'any'.
*
* @param javaType (JavaType): A JavaType object representing the Java type to be mapped.
* @param paramsMap (map[string]string): A map used for generating TypeScript code.
* @param className (string): The name of the current class.
* @return (string): The TypeScript type corresponding to the provided Java type.
*/
func FindTsType(javaType JavaType, paramsMap map[string]string, className string) string {
    tsType, ok := typeMapping[javaType.Name.Value]
    if ok {
        return tsType
    } else if strings.HasSuffix(javaType.Name.Value, config.CONFIG.DtoPackage.Suffix) {
        tsTypeName := utils.RemoveSuffix(javaType.Name.Value, config.CONFIG.DtoPackage.Suffix)
        if tsTypeName != className {
            paramsMap["{%file_import%}"] = utils.RemoveSuffix(utils.ToInterfaceFileName(tsTypeName), ".ts")
            paramsMap["{%new_import%}"] = tsTypeName
            paramsMap["{%imports%}"] += utils.FormatString(paramsMap, angular.ENTITY_IMPORT_TEMPLATE)
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
