package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateService(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass { 
    service := entities.BaseJavaClass{
        Packages : paramsMap["{%service_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaService.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaService.Annotations),
        ClassType : java.JavaService.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%service_suffix%}"],
        Implements : java.JavaService.Implements,
        Extends : java.JavaService.Extends,
        Body : utils.FormatString(paramsMap, java.JavaService.Body),
    }
    service.Directory = findDirectoryPath(service)
    service.FileName = service.ClassName + service.ClassSuffix + ".java"  
    return service
} 
