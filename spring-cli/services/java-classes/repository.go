package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateRepository(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    repository := entities.BaseJavaClass{
        Packages : paramsMap["{%repository_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaRepository.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaRepository.Annotations),
        ClassType : java.JavaRepository.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%repository_suffix%}"],
        Implements : java.JavaRepository.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaRepository.Extends),
        Body : utils.FormatString(paramsMap, java.JavaRepository.Body),
    }
    repository.Directory = findDirectoryPath(repository)
    repository.FileName = repository.ClassName + repository.ClassSuffix + ".java"  
    return repository
} 
