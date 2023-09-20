package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

/* CreateEntity 
*  generates a Java entity class based on a JpaEntity and a parameter map.
*
*  It returns an entities.BaseJavaClass representing the generated Java entity.
*
*  The function creates a simple class using 'CreateSimpleClass' and sets the class type to 'java.JavaEntity'.
*  It then generates the body map with fields using 'createEntityBody'.
*  The 'Body' field of the entity is formatted with the generated body map.
*  Special imports are also formatted based on the parameters map.
*  The resulting entity is returned.
*
*  @param class (entities.JpaEntity): The JpaEntity used to generate the entity class.
*  @param paramsMap (map[string]string): A parameter map containing additional information.
*  @return (entities.BaseJavaClass): A Java entity class generated from the JpaEntity.
*/
func CreateEntity(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    entity := CreateSimpleClass(class, paramsMap, java.JavaEntity)
    bodyMap := map[string]string{
        "{%fields%}" : createEntityBody(class, &entity),
    }
    entity.Body = utils.FormatString(bodyMap, java.JavaEntity.Body)
    entity.SpecialImports = utils.FormatString(paramsMap, java.JavaEntity.SpecialImports)
    return entity
} 


/* createEntityBody 
*  generates the body of a Java entity class based on a JpaEntity and a parameter map.
*
*  The function iterates through the fields of the JpaEntity and updates the imports in the entity.
*  It then creates class fields using 'createClassField' and appends them to the 'fields' slice.
*  Finally, it joins the fields into a string and returns it.
*
*  @param class (entities.JpaEntity): The JpaEntity used to generate the entity body.
*  @param entity (*entities.BaseJavaClass): A pointer to the base Java class entity.
*  @return (string): The generated body of the Java entity as a string.
*/
func createEntityBody(class entities.JpaEntity, entity *entities.BaseJavaClass) string{
    var fields = []string{}
    for _, field := range class.Fields{
        updateImport(field, entity)
        fields = append(fields,createClassField(field, field.Options.Annotations, entity))
    }
    return strings.Join(fields, "")
}

