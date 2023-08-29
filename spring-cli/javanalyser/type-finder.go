package javanalyser

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
func FindTsType(javaType JavaType) string {
    tsType, ok := typeMapping[javaType.Name.Value]
    if ok {
        return tsType
    } else {
        switch javaType.Name.Value {
        case "List":
            return FindTsType(javaType.SubTypes[0]) + "[]"
        case "Set":
            return FindTsType(javaType.SubTypes[0]) + "[]"
        case "ResponseEntity":
            return FindTsType(javaType.SubTypes[0])
        default:
            return "any"
        }
    }
}
