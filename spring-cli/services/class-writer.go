package services

import (

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/services/java-classes"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func createJavaFileBytes(javaClass entities.BaseJavaClass) []byte{
    paramsMap := map[string]string {
        "{%package%}" : javaClass.Packages,
        "{%imports%}" : javaClass.Imports + javaClass.SpecialImports,
        "{%annotations%}" : javaClass.Annotations,
        "{%class_type%}" : javaClass.ClassType,
        "{%class_name%}" : javaClass.ClassName,
        "{%suffix%}" : javaClass.ClassSuffix,
        "{%extends%}" : javaClass.Extends,
        "{%implements%}" : javaClass.Implements,
        "{%body%}" : javaClass.Body,
    }
    return []byte(utils.FormatString(paramsMap, java.ClassTemplate))
}

func writeClassesForOneEntity(classes []entities.BaseJavaClass){
    for _, class := range classes{
        daos.WriteSimpleFile(class.Directory, class.FileName, createJavaFileBytes(class))        
    }
}


func CreateJavaClasses(){
    classes := daos.LoadEntityJson();
    javaclasses.WriteClassImports(classes)
    loadFieldsFromTemplate(&java.JavaController, "Controller/")
    loadFieldsFromTemplate(&java.JavaService, "Service/")
    loadFieldsFromTemplate(&java.JavaRepository, "Repository/")
    loadFieldsFromTemplate(&java.JavaEntity, "Entity/")
    loadFieldsFromTemplate(&java.JavaDto, "Dto/")
    loadFieldsFromTemplate(&java.JavaMapper, "Mapper/")
    for _, class := range classes{
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

func initializeClassesFromTemplates(){
    daos.ReadTemplateFile([]entities.BaseJavaClass{
        java.JavaException,
        java.JavaAnnotation,
        java.JavaInterface,
        java.JavaEnum,
        java.JavaRecord,
        java.JavaBaseClass,
        java.JavaController, 
        java.JavaService, 
        java.JavaDto, 
        java.JavaMapper, 
        java.JavaEntity,
        java.JavaRepository, 
    })
}

func CreateJavaClass(cname, classType string){
    cname, pname := utils.GetClassNameAndPackageFromArgs(cname)
    initializeClassesFromTemplates()
    classInfos := entities.JpaEntity{
        Name : cname,
        Package: config.CONFIG.BasePackage + pname,
    }
    var entity entities.BaseJavaClass
    switch classType{
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
    default :
    }
    daos.WriteSimpleFile(entity.Directory, entity.FileName, createJavaFileBytes(entity))
}

func loadFieldsFromTemplate(javaClass *entities.BaseJavaClass, directoryName string){
    javaClass.Packages = daos.ReadTemplateField(directoryName + "Packages.txt")
    javaClass.Imports = daos.ReadTemplateField(directoryName + "Imports.txt")
    javaClass.SpecialImports = daos.ReadTemplateField(directoryName + "SpecialImports.txt")
    javaClass.Annotations = daos.ReadTemplateField(directoryName + "Annotations.txt")
    javaClass.ClassType = daos.ReadTemplateField(directoryName + "ClassType.txt")
    javaClass.ClassName = daos.ReadTemplateField(directoryName + "ClassName.txt")
    javaClass.ClassSuffix = daos.ReadTemplateField(directoryName + "ClassSuffix.txt")
    javaClass.Implements = daos.ReadTemplateField(directoryName + "Implements.txt")
    javaClass.Extends = daos.ReadTemplateField(directoryName + "Extends.txt")
    javaClass.Body = daos.ReadTemplateField(directoryName + "Body.txt")
    javaClass.FileName = daos.ReadTemplateField(directoryName + "FileName.txt")
    javaClass.Directory = daos.ReadTemplateField(directoryName + "Directory.txt")
}
