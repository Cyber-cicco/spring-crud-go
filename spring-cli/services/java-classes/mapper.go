package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/*CreateMapper 
* generates a Java mapper class based on a JpaEntity and a parameter map.
*
* The function creates a simple class using 'CreateSimpleClass' and updates its special imports.
* It then generates the body of the mapper using 'createMapperBody' and sets it in the mapper class.
* Finally, it returns the generated mapper class.
*
* @param class (entities.JpaEntity): The JpaEntity used to generate the mapper class.
* @param paramsMap (map[string]string): A parameter map used for generating the mapper.
* @return (entities.BaseJavaClass): The generated mapper class.
*/
func CreateMapper(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    mapper := CreateSimpleClass(class, paramsMap, java.JavaMapper)
    mapper.SpecialImports = utils.FormatString(paramsMap, java.JavaMapper.SpecialImports)
    mapper.Body = createMapperBody(class, utils.CopyMap[string, string](paramsMap))    
    return mapper
} 

/* createMapperBody 
*  generates the body of a Java mapper class based on a JpaEntity and a parameter map.
*
*  The function iterates through the fields of the JpaEntity and generates sets for DTO and entity mapping.
*  It uses 'java.MapperSetDto' and 'java.MapperSetEntity' templates to format the sets.
*  The formatted sets are joined into strings and added to the parameter map.
*  Finally, it formats the complete mapper body using the parameter map and returns the result.
*
*  @param object (entities.JpaEntity): The JpaEntity used to generate the mapper body.
*  @param paramsMap (map[string]string): A parameter map used for generating the mapper.
*  @return (string): The generated mapper body.
*/
func createMapperBody(object entities.JpaEntity, paramsMap map[string]string) string{
    var dtoSets = []string{}
    var entitySets = []string{}
    for _, field := range object.Fields{
        mapField := map[string]string { "{%field_title%}" : strings.Title(field.Name) }
        rawType := utils.DenestObject(field.Type)
        _, exists := EntityTypes[rawType]
        if !exists {
            dtoSets = append(dtoSets, utils.FormatString(mapField, java.MapperSetDto))
            entitySets = append(entitySets, utils.FormatString(mapField, java.MapperSetEntity))
        }
    }
    paramsMap["{%sets_dto%}"] = strings.Join(dtoSets, "")
    paramsMap["{%sets_entity%}"] = strings.Join(entitySets, "")
    return utils.FormatString(paramsMap, java.JavaMapper.Body)
}
