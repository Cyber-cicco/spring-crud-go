package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateMapper(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    mapper := entities.BaseJavaClass{
        Packages : paramsMap["{%mapper_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaMapper.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaMapper.Annotations),
        ClassType : java.JavaMapper.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%mapper_suffix%}"],
        Implements : java.JavaMapper.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaMapper.Extends),
    }
    mapper.Body = createMapperBody(class, utils.CopyMap[string, string](paramsMap))    
    mapper.Directory = findDirectoryPath(mapper)
    mapper.FileName = mapper.ClassName + mapper.ClassSuffix + ".java"  
    return mapper
} 

func createMapperBody(object entities.JpaEntity, paramsMap map[string]string) string{
    var dtoSets = []string{}
    var entitySets = []string{}
    for _, field := range object.Fields{
        mapField := map[string]string { "{%field_title%}" : utils.ToTitle(field.Name) }
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
