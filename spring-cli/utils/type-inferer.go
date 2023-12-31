package utils

import (
	"errors"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
)

func InfereTypeByName(fieldName string) string {
    returnedType := "{returned}"

    if strings.Contains(fieldName, ":") {
        fields := strings.Split(fieldName, ":")
        if len(fields) == 1{
            panic(errors.New(config.ERR_TYPE_NOT_GIVEN))
        }
        return fields[1]
    }
    
    if strings.Contains(fieldName, "List") {
        returnedType = "List<{returned}>"
    }

    if strings.Contains(fieldName, "Set") {
        returnedType = "Set<{returned}>"
    }

    if strings.Contains(fieldName, "*"){
        fieldName = strings.ReplaceAll(fieldName, "List", "")
        fieldName = strings.ReplaceAll(fieldName, "Set", "")
        fieldName = strings.ReplaceAll(fieldName, "*", "")
        fieldName = strings.Title(fieldName)
        return strings.Replace(returnedType, "{returned}", fieldName, 1);
    }

    if strings.HasPrefix(fieldName, "id"){
        return strings.Replace(returnedType, "{returned}", "Long", 1);
    }

    if strings.HasPrefix(fieldName, "is"){
        return strings.Replace(returnedType, "{returned}", "Boolean", 1);
    }

    if strings.Contains(fieldName, "Date") || strings.HasPrefix(fieldName, "date"){
        return strings.Replace(returnedType, "{returned}", "LocalDate", 1);
    }

    if strings.Contains(fieldName, "duree") || strings.Contains(fieldName, "longueur") || strings.Contains(fieldName, "length") || strings.Contains(fieldName, "quantit") || strings.HasPrefix(fieldName, "nb") {
        return strings.Replace(returnedType, "{returned}", "Integer", 1);
    }

    if strings.Contains(fieldName, "price") || strings.Contains(fieldName, "prix"){
        return strings.Replace(returnedType, "{returned}", "Double", 1);
    }

    return strings.Replace(returnedType, "{returned}", "String", 1);
}
