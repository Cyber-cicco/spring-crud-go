package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateDto(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    dto := entities.BaseJavaClass{
        Packages : paramsMap["{%dto_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaDto.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaDto.Annotations),
        ClassType : java.JavaDto.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%dto_suffix%}"],
        Implements : java.JavaDto.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaDto.Extends),
    }
    bodyMap := map[string]string{
        "{%fields%}" : createDtoBody(class, &dto),
    }
    dto.Body = utils.FormatString(bodyMap, java.JavaDto.Body)
    dto.Directory = findDirectoryPath(dto)
    dto.FileName = dto.ClassName + dto.ClassSuffix + ".java"  
    return dto
} 

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
        fields = append(fields, createClassField(field))
    }
    return  strings.Join(fields, "")
}
