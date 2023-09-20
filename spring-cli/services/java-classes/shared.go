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

/* updateImport 
*  updates the import statements in a Java class based on the field type.
*
*  It adds import statements for List, Set, and LocalDate if they are not already present in the class.
*
*  @param field (entities.JpaField): The field whose type will be used for updating the imports.
*  @param entity (*entities.BaseJavaClass): A pointer to the Java class to be updated.
*/
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

/*createClassField 
* generates the Java code for a class field with annotations.
*
* The function constructs a map of parameters for string formatting, including annotations, field type, and field name.
* It then uses utils.FormatString to generate the Java code for the field with annotations.
*
* @param field (entities.JpaField): The field for which the code will be generated.
* @param annotations ([]string): A slice of annotations to be applied to the field.
* @param entity (*entities.BaseJavaClass): A pointer to the Java class to which the field belongs.
* @return (string): The generated Java code for the class field.
*/
func createClassField(field entities.JpaField, annotations []string, entity *entities.BaseJavaClass) string {
	paramsMap := map[string]string{
		"{%annotations%}": createAnnotations(annotations, field.Name, entity.ClassName), 
		"{%field_type%}":  field.Type,
		"{%field_name%}":  field.Name,
	}
	return utils.FormatString(paramsMap, java.JavaEntityField)
}

/*createAnnotations 
* generates Java annotations for a field.
*
* The function modifies 'fieldName' by removing occurrences of "List" and "Set".
* It also converts a class Name to an attribute name.
* The function then iterates through the annotations and processes the "mtm" annotation.
* For "mtm" annotations, it generates a many-to-many table name and formats the annotation accordingly.
* The modified annotations are then joined into a single string and returned.
*
* @param annotations ([]string): A slice of annotations to be applied to the field.
* @param fieldName (string): The name of the field.
* @param className (string): The name of the class to which the field belongs.
* @return (string): Java annotations for the field.
*/
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

/*
* Creates a many to many annotation based on the field name and the class name.
* It checks if the annotation already exists in the map and returns it if it does.
*/
func createManyToManyAnnotation(fieldName, className string) string {
    return checkManyToManyExists(fieldName, className)
}

/*checkManyToManyExists 
* checks if a many-to-many relationship already exists in the MTM_MAP.
* If not, it generates a unique key and value based on the field and class names.
*
* It first creates a unique key using 'createMapKey'. That keys is based on the field and class names.
* It then checks if the key exists in the MTM_MAP.
* If not, it generates a value using the class and field names and stores it in the map.
* Finally, it returns the value.
*
* @param fieldName (string): The name of the field.
* @param className (string): The name of the class.
* @return (string): The value associated with the key in MTM_MAP.
*/
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
/* createPackage
*  generates the package name based on the provided entity and package option.
*
*  It applies the specified package policy to determine the resulting package name.
*
*  @param entity (entities.JpaEntity): A JpaEntity object representing the entity.
*  @param option (config.PackageOption): A configuration option for handling packages.
*  @return (string): The generated package name.
*/
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

/* CreateParamsMapAndIrrigateTemplates 
*  generates a parameter map for template filling and sets package information.
*
*  It creates a map with placeholders and their corresponding values for template filling.
*  Additionally, it sets package information for various Java classes.
*
*  @param entity (entities.JpaEntity): A JpaEntity object representing the entity.
*  @return (map[string]string): A map containing template placeholders and their values.
*/
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

/* CreateSimpleClass 
*  creates a simple Java class based on the provided entity and template class.
*
*  The function performs the following tasks:
*    - Copies various properties from the 'noParamClass' template class.
*    - Formats and sets imports, annotations, class type, class name, class suffix, implements, and extends.
*    - Finds the directory path and sets the file name for the class.
*
*  @param class (entities.JpaEntity): A JpaEntity object representing the entity.
*  @param paramsMap (map[string]string): A map containing template placeholders and their values.
*  @param noParamClass (entities.BaseJavaClass): A template class object without parameters.
*  @return (entities.BaseJavaClass): A BaseJavaClass object representing the created Java class.
*/
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
