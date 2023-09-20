package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/* CreateService 
*  generates a Java service class based on a JpaEntity and a parameter map.
*
*  The function uses the 'CreateSimpleClass' function to initialize the service.
*  It then sets the body and special imports properties using the parameter map.
*  Finally, it returns the generated service.
*
*  @param class (entities.JpaEntity): The JpaEntity used to generate the service.
*  @param paramsMap (map[string]string): A parameter map used for generating the service.
*  @return (entities.BaseJavaClass): The generated service.
*/
func CreateService(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass { 
    service := CreateSimpleClass(class, paramsMap, java.JavaService)
    service.Body = utils.FormatString(paramsMap, java.JavaService.Body)
    service.SpecialImports = utils.FormatString(paramsMap, java.JavaService.SpecialImports)
    return service
} 
