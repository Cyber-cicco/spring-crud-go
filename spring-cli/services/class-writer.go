package services

import (
	"fmt"

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
        daos.WriteJavaClass(class.Directory, class.FileName, createJavaFileBytes(class))        
    }
}


func CreateJavaClasses(){
    classes := daos.LoadEntityJson();
    javaclasses.WriteClassImports(classes)
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

func CreateJavaClass(cname, classType string){
    cname, pname := utils.GetClassNameAndPackageFromArgs(cname)
    classInfos := entities.JpaEntity{
        Name : cname,
        Package: config.CONFIG.BasePackage + pname,
    }
    var entity entities.BaseJavaClass
    switch classType{
    case "ctrl":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaController)
    case "srv":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaService)
    case "ent":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaEntity)
    case "map":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaMapper)
    case "dto":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaDto)
    case "repo":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaRepository)
    case "ex":
        entity = javaclasses.CreateSimpleClass(classInfos, javaclasses.CreateParamsMapAndIrrigateTemplates(classInfos), java.JavaException)
    case "":
    default :
    }
    fmt.Println(entity.Directory)
    daos.WriteJavaClass(entity.Directory, entity.FileName, createJavaFileBytes(entity))
}
