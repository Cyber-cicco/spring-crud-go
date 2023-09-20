package javaclasses

import (

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/* CreateController 
*  generates a Java controller class based on a Jpa config object and parameter map.
*
*  The function uses CreateSimpleClass to create a basic class structure with appropriate annotations.
*  It then updates the special imports and body of the controller using formatted strings from the parameter map.
*
*  @param class (entities.JpaEntity): The JpaEntity for which the controller is being created.
*  @param paramsMap (map[string]string): A map of parameters used for generating the controller class.
*  @return (entities.BaseJavaClass): A BaseJavaClass representing the generated controller.
*/func CreateController(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
	controller := CreateSimpleClass(class, paramsMap, java.JavaController)
	controller.SpecialImports = utils.FormatString(paramsMap, java.JavaController.SpecialImports)
	controller.Body = utils.FormatString(paramsMap, java.JavaController.Body)
	return controller

}
