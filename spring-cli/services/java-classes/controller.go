package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)


func CreateController(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    controller := entities.BaseJavaClass{
        Packages : paramsMap["{%controller_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaController.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaController.Annotations),
        ClassType : java.JavaController.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%controller_suffix%}"],
        Implements : java.JavaController.Implements,
        Extends : java.JavaController.Extends,
        Body : utils.FormatString(paramsMap, java.JavaController.Body),
    }
    controller.Directory = findDirectoryPath(controller)
    controller.FileName = controller.ClassName + controller.ClassSuffix + ".java"  
    return controller
   
}
