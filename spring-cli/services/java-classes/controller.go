package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateController(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
	controller := CreateSimpleClass(class, paramsMap, java.JavaController)
	controller.SpecialImports = utils.FormatString(paramsMap, java.JavaController.SpecialImports)
	controller.Body = utils.FormatString(paramsMap, java.JavaController.Body)
	return controller

}
