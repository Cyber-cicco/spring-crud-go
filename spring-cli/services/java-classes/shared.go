package javaclasses

import (
	"fmt"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

var MANY_TO_MANY = `@ManyToMany
    @JoinTable(name = "{%mtm_table_name%}",
            joinColumns = @JoinColumn(name = "{%class_name%}_id", referencedColumnName = "id"),
            inverseJoinColumns = @JoinColumn(name = "{%target_class_name%}_id", referencedColumnName = "id")
    )
    `

var MTM_MAP = map[uint64]string{}

var EntityTypes = map[string]string{}

func findDirectoryPath(javaClass entities.BaseJavaClass) string {
	return config.CONFIG.BaseJavaDir + strings.ReplaceAll(javaClass.Packages, ".", "/") + "/"
}

func updateImport(field entities.JpaField, entity *entities.BaseJavaClass) {
	if strings.Contains(field.Type, "List<") && !strings.Contains(entity.Imports, ".List;"){
		entity.Imports += "\nimport java.util.List;"
	}
	if strings.Contains(field.Type, "Set<") && !strings.Contains(entity.Imports, ".Set;"){
		entity.Imports += "\nimport java.util.Set;"
	}
    if strings.Contains(field.Type, "LocalDate") && !strings.Contains(entity.Imports, ".LocalDate;"){
		entity.Imports += "\nimport java.time.LocalDate;"
    }
}

func createClassField(field entities.JpaField, annotations []string, entity *entities.BaseJavaClass) string {
	paramsMap := map[string]string{
		"{%annotations%}": createAnnotations(annotations, field.Name, entity.ClassName), 
		"{%field_type%}":  field.Type,
		"{%field_name%}":  field.Name,
	}
	return utils.FormatString(paramsMap, java.JavaEntityField)
}
func createAnnotations(annotations []string, fieldName, className string) string {
	fieldName = strings.ReplaceAll(fieldName, "List", "")
	fieldName = strings.ReplaceAll(fieldName, "Set", "")
    attributeName := utils.ToAttributeName(className)
    for i, annotation := range annotations {
        if annotation == "mtm" {
            mtmTableName := createManyToManyAnnotation(fieldName, attributeName)
            paramsMap := map[string]string{
                "{%mtm_table_name%}": mtmTableName,
                "{%target_class_name%}": fieldName,
                "{%class_name%}": attributeName,
            }
            annotations[i] = utils.FormatString(paramsMap, MANY_TO_MANY)
            fmt.Printf("annotations[i]: %v\n", annotations[i])
        }
    }
    return strings.Join(annotations, "")
}
func createManyToManyAnnotation(fieldName, className string) string {
    return checkManyToManyExists(fieldName, className)
}
func checkManyToManyExists(fieldName, className string) string {
	key := createMapKey(fieldName, className)
	val, ok := MTM_MAP[key]
	if !ok {
		val = className + "_" + fieldName
        MTM_MAP[key] = val
	}
	return val
}

/**
* prend le nom du field et du nom de la class pour en trouver une clé en faisant un XOR dessus.
* Comme ça, dans le cas d'un autre many to many désignant la même solution mais dans l'autre classe
* cela permettra de récupérer le many to many déjà créé.
 */
func createMapKey(fieldName, className string) uint64 {
	var sum uint64
	for _, val := range fieldName {
		sum += uint64(val)
	}
	for _, val := range className {
		sum += uint64(val)
	}
	return sum 
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
		"{%class_name%}":         strings.Title(entity.Name),
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
