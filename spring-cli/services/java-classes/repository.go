package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateRepository(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    repository := CreateSimpleClass(class, paramsMap, java.JavaRepository)
    repository.SpecialImports = utils.FormatString(paramsMap, java.JavaRepository.SpecialImports)
    repository.Body = utils.FormatString(paramsMap, java.JavaRepository.Body)
    repository.Extends = utils.FormatString(paramsMap, java.JavaRepository.Extends)
    return repository
} 
