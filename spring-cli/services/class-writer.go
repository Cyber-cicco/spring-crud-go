package services

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
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
        "{%class_name%}" : utils.ToClassName(entity.Name),
        "{%class_name_lower%}" : utils.ToAttributeName(entity.Name),

    }
    return paramsMap
}

func createController(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    controller := entities.BaseJavaClass{
        Packages : paramsMap["{%controller_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaController.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaController.Annotations),
        ClassType : java.JavaController.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%controller_suffix%}"],
        Implements : java.JavaController.Implements,
        Extends : java.JavaController.Extends,
        Body : utils.FormatString(paramsMap, java.JavaController.Body),
    }
    controller.Directory = findDirectoryPath(controller)
    controller.FileName = controller.ClassName + controller.ClassSuffix + ".java"  
    return controller
   
}
func createService(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass { 
    service := entities.BaseJavaClass{
        Packages : paramsMap["{%service_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaService.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaService.Annotations),
        ClassType : java.JavaService.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%service_suffix%}"],
        Implements : java.JavaService.Implements,
        Extends : java.JavaService.Extends,
        Body : utils.FormatString(paramsMap, java.JavaService.Body),
    }
    service.Directory = findDirectoryPath(service)
    service.FileName = service.ClassName + service.ClassSuffix + ".java"  
    return service
} 

func createRepository(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    repository := entities.BaseJavaClass{
        Packages : paramsMap["{%repository_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaRepository.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaRepository.Annotations),
        ClassType : java.JavaRepository.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%repository_suffix%}"],
        Implements : java.JavaRepository.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaRepository.Extends),
        Body : utils.FormatString(paramsMap, java.JavaRepository.Body),
    }
    repository.Directory = findDirectoryPath(repository)
    repository.FileName = repository.ClassName + repository.ClassSuffix + ".java"  
    return repository
} 


func createJavaFileBytes(javaClass entities.BaseJavaClass) []byte{
    paramsMap := map[string]string {
        "{%package%}" : javaClass.Packages,
        "{%imports%}" : javaClass.Imports,
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

func findDirectoryPath(javaClass entities.BaseJavaClass) string{
    return config.CONFIG.BaseJavaDir + strings.ReplaceAll(javaClass.Packages, ".", "/") + "/" 
}

func writeClassesForOneEntity(classes []entities.BaseJavaClass){
    for _, class := range classes{
        daos.WriteJavaClass(class.Directory, class.FileName, createJavaFileBytes(class))        
    }
}

func CreateJavaClasses(){
    classes, err := daos.LoadEntityJson();
    utils.HandleBasicError(err, "Il n'y a pas de fichier de configuration d'entités JPA dans le dossier jpa. Veuillez générer les fichier de configuration à l'aide de la commande jpa avant d'utiliser la commande spring")
    for _, class := range classes{
        paramsMap := createParamsMap(class)
        classList := []entities.BaseJavaClass{
             createController(class, paramsMap),
             createService(class, paramsMap),
             createRepository(class, paramsMap),
        }
        writeClassesForOneEntity(classList)
    }
}
