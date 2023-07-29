package utils

import (
	"errors"
	"strings"
)

func InfereTypeByName(fieldName string) string {
    returnedType := "{returned}"

    if strings.Contains(fieldName, ":") {
        fields := strings.Split(fieldName, ":")
        if len(fields) == 1{
            panic(errors.New("Le caractère ':' est utilisé pour préciser le type. Ce qui suit ce caractère ne doit pas être vide"))
        }
        return fields[1]
    }
    
    if strings.Contains(fieldName, "List") || string(fieldName[len(fieldName)-1]) == "s"{
        returnedType = "List<{returned}>"
    }

    if strings.Contains(fieldName, "Set") {
        returnedType = "Set<{returned}>"
    }

    if strings.Contains(fieldName, "@"){
        fieldName = strings.ReplaceAll(fieldName, "List", "")
        fieldName = strings.ReplaceAll(fieldName, "@", "")
        fieldName = ToTitle(fieldName)
        return strings.Replace(returnedType, "{returned}", fieldName, 1);
    }

    if strings.HasPrefix(fieldName, "id"){
        return strings.Replace(returnedType, "{returned}", "Long", 1);
    }

    if strings.HasPrefix(fieldName, "is"){
        return strings.Replace(returnedType, "{returned}", "Boolean", 1);
    }

    if strings.Contains(fieldName, "Date") || strings.HasPrefix(fieldName, "date"){
        return strings.Replace(returnedType, "{returned}", "Date", 1);
    }

    if strings.Contains(fieldName, "duree") || strings.Contains(fieldName, "longueur") || strings.Contains(fieldName, "length") || strings.Contains(fieldName, "quantit") || strings.HasPrefix(fieldName, "nb") {
        return strings.Replace(returnedType, "{returned}", "Integer", 1);
    }

    if strings.Contains(fieldName, "price") || strings.Contains(fieldName, "prix"){
        return strings.Replace(returnedType, "{returned}", "Double", 1);
    }

    if strings.Contains(strings.ToLower(fieldName), "nom") || strings.Contains(strings.ToLower(fieldName), "name") {
        return strings.Replace(returnedType, "{returned}", "Double", 1);
    }

    return strings.Replace(returnedType, "{returned}", "String", 1);
}
