package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateEntity(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    entity := entities.BaseJavaClass{
        Packages : paramsMap["{%entity_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaEntity.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaEntity.Annotations),
        ClassType : java.JavaEntity.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%entity_suffix%}"],
        Implements : java.JavaEntity.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaEntity.Extends),
    }
    bodyMap := map[string]string{
        "{%fields%}" : createEntityBody(class, &entity),
    }
    entity.Body = utils.FormatString(bodyMap, java.JavaEntity.Body)
    entity.Directory = findDirectoryPath(entity)
    entity.FileName = entity.ClassName + entity.ClassSuffix + ".java"  
    return entity
} 

func createEntityBody(class entities.JpaEntity, entity *entities.BaseJavaClass) string{
    var fields = []string{}
    for _, field := range class.Fields{
        updateImport(field, entity)
        fields = append(fields,createClassField(field, field.Options.Annotations))
    }
    return strings.Join(fields, "")
}

