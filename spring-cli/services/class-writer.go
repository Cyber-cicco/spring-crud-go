package services

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)

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

func createParamsMap(entity entities.JpaEntity) map[string]string{
    paramsMap := map[string]string {
        "{dto_package}" : createPackage(entity, config.CONFIG.DtoPackage),
        "{mapper_package}" : createPackage(entity, config.CONFIG.MapperPackage),
        "{repository_package}" : createPackage(entity, config.CONFIG.RepositoryPackage),
        "{entity_package}" : createPackage(entity, config.CONFIG.EntityPackage),
        "{service_package}" : createPackage(entity, config.CONFIG.ServicePackage),
        "{controller_package}" : createPackage(entity, config.CONFIG.ControllerPackage),
        "{dto_suffix}" : createPackage(entity, config.CONFIG.DtoPackage),
        "{mapper_suffix}" :  config.CONFIG.MapperPackage.Suffix,
        "{repository_suffix}" :  config.CONFIG.RepositoryPackage.Suffix,
        "{entity_suffix}" :  config.CONFIG.EntityPackage.Suffix,
        "{service_suffix}" : config.CONFIG.ServicePackage.Suffix,
        "{controller_suffix}" :  config.CONFIG.ControllerPackage.Suffix,
        "{class_name}" : utils.ToClassName(entity.Name),
        "{class_name_lower}" : utils.ToAttributeName(entity.Name),

    }
    return paramsMap
}

func CreateJavaClasses(){
    classes := daos.LoadEntityJson();
    fmt.Println(classes)
}
