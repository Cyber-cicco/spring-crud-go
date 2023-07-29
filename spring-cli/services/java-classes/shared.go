package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

var EntityTypes = map[string]string{}

func findDirectoryPath(javaClass entities.BaseJavaClass) string{
    return config.CONFIG.BaseJavaDir + strings.ReplaceAll(javaClass.Packages, ".", "/") + "/" 
}

func updateImport(field entities.JpaField, entity *entities.BaseJavaClass){
    if strings.Contains(field.Type, "List<"){
        entity.Imports += "\nimport java.util.List;"
    }
    if strings.Contains(field.Type, "Set<"){
        entity.Imports += "\nimport java.util.Set;"
    }
}

func createClassField(field entities.JpaField) string {
    paramsMap := map[string]string {
        "{%annotations%}":"",
        "{%field_type%}":field.Type,
        "{%field_name%}":field.Name,
    }
    return utils.FormatString(paramsMap, java.JavaEntityField)
}

func WriteClassImports(classes []entities.JpaEntity){
    for _, class := range classes{
        EntityTypes[class.Name] = createPackage(class, config.CONFIG.DtoPackage)
    }
}

func createPackage(entity entities.JpaEntity,  option config.PackageOption) string{
    if option.PackagePolicy == "appended" {
        if option.Package == "" {
            return entity.Package 
        }
        return entity.Package + "." + option.Package
    }
    if option.PackagePolicy == "base" && option.Package != "" {
        return config.CONFIG.BasePackage + "." + option.Package
    }
    return config.CONFIG.BasePackage
}

func CreateParamsMap(entity entities.JpaEntity) map[string]string{
    paramsMap := map[string]string {
        "{%dto_package%}" : createPackage(entity, config.CONFIG.DtoPackage),
        "{%mapper_package%}" : createPackage(entity, config.CONFIG.MapperPackage),
        "{%repository_package%}" : createPackage(entity, config.CONFIG.RepositoryPackage),
        "{%entity_package%}" : createPackage(entity, config.CONFIG.EntityPackage),
        "{%service_package%}" : createPackage(entity, config.CONFIG.ServicePackage),
        "{%controller_package%}" : createPackage(entity, config.CONFIG.ControllerPackage),
        "{%dto_suffix%}" : config.CONFIG.DtoPackage.Suffix,
        "{%mapper_suffix%}" :  config.CONFIG.MapperPackage.Suffix,
        "{%repository_suffix%}" :  config.CONFIG.RepositoryPackage.Suffix,
        "{%entity_suffix%}" :  config.CONFIG.EntityPackage.Suffix,
        "{%service_suffix%}" : config.CONFIG.ServicePackage.Suffix,
        "{%controller_suffix%}" :  config.CONFIG.ControllerPackage.Suffix,
        "{%class_name%}" : utils.ToTitle(entity.Name),
        "{%class_name_lower%}" : utils.ToAttributeName(entity.Name),

    }
    return paramsMap
}
