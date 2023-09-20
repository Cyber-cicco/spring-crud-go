package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateEntity(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    entity := CreateSimpleClass(class, paramsMap, java.JavaEntity)
    bodyMap := map[string]string{
        "{%fields%}" : createEntityBody(class, &entity),
    }
    entity.Body = utils.FormatString(bodyMap, java.JavaEntity.Body)
    entity.SpecialImports = utils.FormatString(paramsMap, java.JavaEntity.SpecialImports)
    return entity
} 

func createEntityBody(class entities.JpaEntity, entity *entities.BaseJavaClass) string{
    var fields = []string{}
    for _, field := range class.Fields{
        updateImport(field, entity)
        fields = append(fields,createClassField(field, field.Options.Annotations, entity))
    }
    return strings.Join(fields, "")
}

