package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateDto(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    dto := CreateSimpleClass(class, paramsMap, java.JavaDto)
    bodyMap := map[string]string{
        "{%fields%}" : createDtoBody(class, &dto),
    }
    dto.SpecialImports = utils.FormatString(paramsMap, java.JavaDto.SpecialImports)
    dto.Body = utils.FormatString(bodyMap, java.JavaDto.Body)
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
        fields = append(fields, createClassField(field, []string{}))
    }
    return  strings.Join(fields, "")
}

