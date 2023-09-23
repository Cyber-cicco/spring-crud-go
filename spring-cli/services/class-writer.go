package services

import (
    "strings"

    "fr.cybercicco/springgo/spring-cli/config"
    "fr.cybercicco/springgo/spring-cli/daos"
    "fr.cybercicco/springgo/spring-cli/entities"
    javaclasses "fr.cybercicco/springgo/spring-cli/services/java-classes"
    "fr.cybercicco/springgo/spring-cli/templates/java"
    "fr.cybercicco/springgo/spring-cli/utils"
)

/*
* createJavaFileBytes generates the byte representation of a Java class based on the provided BaseJavaClass object.
*
* This function takes one parameter:
*   - 'javaClass' (entities.BaseJavaClass): A BaseJavaClass object representing the Java class.
*
* It returns a slice of bytes containing the byte representation of the Java class.
*
* The function performs the following tasks:
*   - Initializes a paramsMap with placeholders and their corresponding values based on 'javaClass' properties.
*   - Formats and retrieves the Java class template using the provided paramsMap.
*   - Converts the formatted Java class template into a slice of bytes and returns it.
*
@param javaClass (entities.BaseJavaClass): A BaseJavaClass object representing the Java class.
* @return ([]byte): A slice of bytes containing the byte representation of the Java class.
*/
func createJavaFileBytes(javaClass entities.BaseJavaClass) []byte {
    paramsMap := map[string]string{
        "{%package%}":     javaClass.Packages,
        "{%imports%}":     javaClass.Imports + javaClass.SpecialImports,
        "{%annotations%}": javaClass.Annotations,
        "{%class_type%}":  javaClass.ClassType,
        "{%class_name%}":  javaClass.ClassName,
        "{%suffix%}":      javaClass.ClassSuffix,
        "{%extends%}":     javaClass.Extends,
        "{%implements%}":  javaClass.Implements,
        "{%body%}":        javaClass.Body,
    }
    return []byte(utils.FormatString(paramsMap, java.ClassTemplate))
}

func writeClassesForOneEntity(classes []entities.BaseJavaClass) {
    for _, class := range classes {
        daos.WriteSimpleFile(class.Directory, class.FileName, createJavaFileBytes(class))
    }
}

func CreateJavaClasses() {
    classes := daos.LoadEntityJson()
    javaclasses.WriteClassImports(classes)
    loadFieldsFromTemplate(&java.JavaController, "Controller/")
    loadFieldsFromTemplate(&java.JavaService, "Service/")
    loadFieldsFromTemplate(&java.JavaRepository, "Repository/")
    loadFieldsFromTemplate(&java.JavaEntity, "Entity/")
    loadFieldsFromTemplate(&java.JavaDto, "Dto/")
    loadFieldsFromTemplate(&java.JavaMapper, "Mapper/")
    for _, class := range classes {
        paramsMap := javaclasses.CreateParamsMapAndIrrigateTemplates(class)
        classList := []entities.BaseJavaClass{
            javaclasses.CreateController(class, paramsMap),
            javaclasses.CreateService(class, paramsMap),
            javaclasses.CreateRepository(class, paramsMap),
            javaclasses.CreateEntity(class, paramsMap),
            javaclasses.CreateDto(class, paramsMap),
            javaclasses.CreateMapper(class, paramsMap),
        }
        writeClassesForOneEntity(classList)
    }
}

func CreateJavaClass(cname, classTypes string) {
    cnames := strings.Split(cname, " ")
    classTypesArray := strings.Split(classTypes, " ")
    for _, className := range cnames {
        cname, pname := utils.GetClassNameAndPackageFromArgs(className)
        classInfos := entities.JpaEntity{
            Name:    cname,
            Package: config.CONFIG.BasePackage + pname,
        }
        var entity entities.BaseJavaClass
        for _, classType := range classTypesArray {
            switch classType {
            case "ctrl":
                loadFieldsFromTemplate(&java.JavaController, "Controller/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaController)
            case "srv":
                loadFieldsFromTemplate(&java.JavaService, "Service/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaService)
            case "ent":
                loadFieldsFromTemplate(&java.JavaEntity, "Entity/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaEntity)
            case "map":
                loadFieldsFromTemplate(&java.JavaMapper, "Mapper/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaMapper)
            case "dto":
                loadFieldsFromTemplate(&java.JavaDto, "Dto/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaDto)
            case "repo":
                loadFieldsFromTemplate(&java.JavaRepository, "Repository/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaRepository)
            case "ex":
                loadFieldsFromTemplate(&java.JavaException, "Exception/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaException)
            case "enum":
                loadFieldsFromTemplate(&java.JavaEnum, "Enum/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaEnum)
            case "int":
                loadFieldsFromTemplate(&java.JavaInterface, "Interface/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaInterface)
            case "rec":
                loadFieldsFromTemplate(&java.JavaRecord, "Record/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaRecord)
                entity.ClassSuffix = "()"
            case "ano":
                loadFieldsFromTemplate(&java.JavaAnnotation, "Annotation/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaAnnotation)
            case "":
                loadFieldsFromTemplate(&java.JavaBaseClass, "BaseClass/")
                entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaBaseClass)
            default:
            }
            daos.WriteSimpleFile(entity.Directory, entity.FileName, createJavaFileBytes(entity))
        }
    }
}

func loadFieldsFromTemplate(javaClass *entities.BaseJavaClass, directoryName string) {
    javaClass.Imports = daos.ReadTemplateField(directoryName + "Imports.txt")
    javaClass.Annotations = daos.ReadTemplateField(directoryName + "Annotations.txt")
    javaClass.ClassType = daos.ReadTemplateField(directoryName + "ClassType.txt")
    javaClass.Implements = daos.ReadTemplateField(directoryName + "Implements.txt")
    javaClass.Extends = daos.ReadTemplateField(directoryName + "Extends.txt")
    javaClass.Body = daos.ReadTemplateField(directoryName + "Body.txt")
}
