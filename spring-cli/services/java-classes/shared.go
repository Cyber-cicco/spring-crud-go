package javaclasses

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

var EntityTypes = map[string]string{}

func findDirectoryPath(javaClass entities.BaseJavaClass) string {
	return config.CONFIG.BaseJavaDir + strings.ReplaceAll(javaClass.Packages, ".", "/") + "/"
}

func updateImport(field entities.JpaField, entity *entities.BaseJavaClass) {
	if strings.Contains(field.Type, "List<") {
		entity.Imports += "\nimport java.util.List;"
	}
	if strings.Contains(field.Type, "Set<") {
		entity.Imports += "\nimport java.util.Set;"
	}
}

func createClassField(field entities.JpaField, annotations []string) string {
	paramsMap := map[string]string{
		"{%annotations%}": strings.Join(annotations, ""),
		"{%field_type%}":  field.Type,
		"{%field_name%}":  field.Name,
	}
	return utils.FormatString(paramsMap, java.JavaEntityField)
}

func WriteClassImports(classes []entities.JpaEntity) {
	for _, class := range classes {
		EntityTypes[class.Name] = createPackage(class, config.CONFIG.DtoPackage)
	}
}

func createPackage(entity entities.JpaEntity, option config.PackageOption) string {
	if option.PackagePolicy == "appended" {
		if option.Package == "" {
			return entity.Package
		}
		return entity.Package + "." + option.Package
	}
	if option.PackagePolicy == "base" && option.Package != "" {
		return entity.Package + "." + option.Package
	}
	return entity.Package
}

func CreateParamsMapAndIrrigateTemplates(entity entities.JpaEntity) map[string]string {
	paramsMap := map[string]string{
		"{%dto_package%}":        createPackage(entity, config.CONFIG.DtoPackage),
		"{%mapper_package%}":     createPackage(entity, config.CONFIG.MapperPackage),
		"{%repository_package%}": createPackage(entity, config.CONFIG.RepositoryPackage),
		"{%entity_package%}":     createPackage(entity, config.CONFIG.EntityPackage),
		"{%service_package%}":    createPackage(entity, config.CONFIG.ServicePackage),
		"{%controller_package%}": createPackage(entity, config.CONFIG.ControllerPackage),
		"{%exception_package%}":  createPackage(entity, config.CONFIG.ExceptionPackage),
		"{%base_package%}":       createPackage(entity, config.CONFIG.DefaultPackage),
		"{%dto_suffix%}":         config.CONFIG.DtoPackage.Suffix,
		"{%mapper_suffix%}":      config.CONFIG.MapperPackage.Suffix,
		"{%repository_suffix%}":  config.CONFIG.RepositoryPackage.Suffix,
		"{%entity_suffix%}":      config.CONFIG.EntityPackage.Suffix,
		"{%service_suffix%}":     config.CONFIG.ServicePackage.Suffix,
		"{%controller_suffix%}":  config.CONFIG.ControllerPackage.Suffix,
		"{%exception_suffix%}":   config.CONFIG.ExceptionPackage.Suffix,
		"{%base_suffix%}":        config.CONFIG.DefaultPackage.Suffix,
		"{%class_name%}":         utils.ToTitle(entity.Name),
		"{%class_name_lower%}":   utils.ToAttributeName(entity.Name),
	}
	java.JavaController.Packages = paramsMap["{%controller_package%}"]
	java.JavaDto.Packages = paramsMap["{%dto_package%}"]
	java.JavaMapper.Packages = paramsMap["{%mapper_package%}"]
	java.JavaRepository.Packages = paramsMap["{%repository_package%}"]
	java.JavaService.Packages = paramsMap["{%service_package%}"]
	java.JavaEntity.Packages = paramsMap["{%entity_package%}"]
	java.JavaException.Packages = paramsMap["{%exception_package%}"]
	java.JavaBaseClass.Packages = paramsMap["{%base_package%}"]
	java.JavaEnum.Packages = paramsMap["{%base_package%}"]
	java.JavaInterface.Packages = paramsMap["{%base_package%}"]
	java.JavaRecord.Packages = paramsMap["{%base_package%}"]
	java.JavaAnnotation.Packages = paramsMap["{%base_package%}"]
	java.JavaController.ClassSuffix = paramsMap["{%controller_suffix%}"]
	java.JavaDto.ClassSuffix = paramsMap["{%dto_suffix%}"]
	java.JavaMapper.ClassSuffix = paramsMap["{%mapper_suffix%}"]
	java.JavaRepository.ClassSuffix = paramsMap["{%repository_suffix%}"]
	java.JavaService.ClassSuffix = paramsMap["{%service_suffix%}"]
	java.JavaEntity.ClassSuffix = paramsMap["{%entity_suffix%}"]
	java.JavaException.ClassSuffix = paramsMap["{%exception_suffix%}"]
	java.JavaBaseClass.ClassSuffix = paramsMap["{%base_suffix%}"]
	java.JavaEnum.ClassSuffix = paramsMap["{%base_suffix%}"]
	java.JavaInterface.ClassSuffix = paramsMap["{%base_suffix%}"]
	java.JavaRecord.ClassSuffix = paramsMap["{%base_suffix%}"]
	java.JavaAnnotation.ClassSuffix = paramsMap["{%base_suffix%}"]
	return paramsMap
}

func CreateSimpleClass(class entities.JpaEntity, paramsMap map[string]string, noParamClass entities.BaseJavaClass) entities.BaseJavaClass {
	paramClass := entities.BaseJavaClass{
		Packages:    noParamClass.Packages,
		Imports:     utils.FormatString(paramsMap, noParamClass.Imports),
		Annotations: utils.FormatString(paramsMap, noParamClass.Annotations),
		ClassType:   noParamClass.ClassType,
		ClassName:   class.Name,
		ClassSuffix: noParamClass.ClassSuffix,
		Implements:  utils.FormatString(paramsMap, noParamClass.Implements),
		Extends:     utils.FormatString(paramsMap, noParamClass.Extends),
	}
	paramClass.Directory = findDirectoryPath(paramClass)
	paramClass.FileName = paramClass.ClassName + paramClass.ClassSuffix + ".java"
	return paramClass
}
