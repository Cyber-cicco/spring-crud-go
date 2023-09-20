package javaclasses

import (
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/* CreateRepository 
*  generates a Java repository class based on a JpaEntity and a parameter map.
*
*  The function uses the 'CreateSimpleClass' function to initialize the repository.
*  It then sets special imports, body, and extends properties using the parameter map.
*  Finally, it returns the generated repository.
*
*  @param class (entities.JpaEntity): The JpaEntity used to generate the repository.
*  @param paramsMap (map[string]string): A parameter map used for generating the repository.
*  @return (entities.BaseJavaClass): The generated repository.
*/
func CreateRepository(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
	repository := CreateSimpleClass(class, paramsMap, java.JavaRepository)
	repository.SpecialImports = utils.FormatString(paramsMap, java.JavaRepository.SpecialImports)
	repository.Body = utils.FormatString(paramsMap, java.JavaRepository.Body)
	repository.Extends = utils.FormatString(paramsMap, java.JavaRepository.Extends)
	return repository
}
