package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateService(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass { 
    service := CreateSimpleClass(class, paramsMap, java.JavaService)
    service.Body = utils.FormatString(paramsMap, java.JavaService.Body)
    service.SpecialImports = utils.FormatString(paramsMap, java.JavaService.SpecialImports)
    return service
} 
