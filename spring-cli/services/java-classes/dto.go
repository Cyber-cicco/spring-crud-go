package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/*CreateDto 
* generates a Java DTO class based on a JpaEntity and parameter map.

* It returns a BaseJavaClass representing the generated DTO.
*
* The function uses CreateSimpleClass to create a basic class structure 
* It then updates the special imports and body of the DTO using formatted strings from the parameter map.
*
* @param class (entities.JpaEntity): The JpaEntity for which the DTO is being created.
* @param paramsMap (map[string]string): A map of parameters used for generating the DTO class.
* @return (entities.BaseJavaClass): A BaseJavaClass representing the generated DTO.
*/
func CreateDto(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    dto := CreateSimpleClass(class, paramsMap, java.JavaDto)
    bodyMap := map[string]string{
        "{%fields%}" : createDtoBody(class, &dto),
    }
    dto.SpecialImports = utils.FormatString(paramsMap, java.JavaDto.SpecialImports)
    dto.Body = utils.FormatString(bodyMap, java.JavaDto.Body)
    return dto
} 

/* createDtoBody 
*  generates the fields of a Java DTO class based on a JpaEntity and updates imports if needed.
*
*  The function iterates through the fields of the JpaEntity and processes them to generate DTO fields.
*  It denests the object type, checks for corresponding imports, and updates them if necessary.
*  It then calls 'updateImport' to handle import statements and appends the generated field to the 'fields' slice.
*  Finally, it joins the fields into a single string and returns it.
*
*  @param object (entities.JpaEntity): The JpaEntity for which the DTO fields are being generated.
*  @param entity (*entities.BaseJavaClass): A pointer to the Java class entity where the fields will be added.
*  @return (string): A string containing the fields of the DTO class.
*/
func createDtoBody(object entities.JpaEntity, entity *entities.BaseJavaClass) string{
    var fields = []string{}
    for _, field := range object.Fields{
        rawType := utils.DenestObject(field.Type)
        javaImport, exists := EntityTypes[rawType]
        if exists {
            entity.Imports += "\nimport " + javaImport + "." + rawType + config.CONFIG.DtoPackage.Suffix + ";"
            field.Type = strings.ReplaceAll(field.Type, rawType, rawType + config.CONFIG.DtoPackage.Suffix)
        }
        updateImport(field, entity)
        fields = append(fields, createClassField(field, []string{}, entity))
    }
    return  strings.Join(fields, "")
}

