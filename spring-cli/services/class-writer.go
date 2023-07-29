package services

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

var basicTypes = []string{
    "String",
    "Long",
    "Integer",
    "Boolean",
    "int",
    "double",
    "Double",
    "Float",
    "float",
    "LocaleDate",
    "LocaleDateTime",
    "LocaleTime",
    "Date",
    "Time",
    "Instant",
}

var EntityTypes = map[string]string{}

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
        "{%class_name%}" : utils.ToTitle(entity.Name),
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

func createEntity(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    entity := entities.BaseJavaClass{
        Packages : paramsMap["{%entity_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaEntity.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaEntity.Annotations),
        ClassType : java.JavaEntity.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%entity_suffix%}"],
        Implements : java.JavaEntity.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaEntity.Extends),
    }
    bodyMap := map[string]string{
        "{%fields%}" : createEntityBody(class, &entity),
    }
    entity.Body = utils.FormatString(bodyMap, java.JavaEntity.Body)
    entity.Directory = findDirectoryPath(entity)
    entity.FileName = entity.ClassName + entity.ClassSuffix + ".java"  
    return entity
} 

func createEntityBody(class entities.JpaEntity, entity *entities.BaseJavaClass) string{
    var fields = []string{}
    for _, field := range class.Fields{
        updateImport(field, entity)
        fields = append(fields, createClassField(field))
    }
    return  strings.Join(fields, "")
}

func updateImport(field entities.JpaField, entity *entities.BaseJavaClass){
    if strings.Contains(field.Type, "List<"){
        entity.Imports += "\nimport java.util.List;"
    }
    if strings.Contains(field.Type, "Set<"){
        entity.Imports += "\nimport java.util.Set;"
    }
}

func createDto(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    dto := entities.BaseJavaClass{
        Packages : paramsMap["{%dto_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaDto.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaDto.Annotations),
        ClassType : java.JavaDto.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%dto_suffix%}"],
        Implements : java.JavaDto.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaDto.Extends),
    }
    bodyMap := map[string]string{
        "{%fields%}" : createDtoBody(class, &dto),
    }
    dto.Body = utils.FormatString(bodyMap, java.JavaDto.Body)
    dto.Directory = findDirectoryPath(dto)
    dto.FileName = dto.ClassName + dto.ClassSuffix + ".java"  
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
        fields = append(fields, createClassField(field))
    }
    return  strings.Join(fields, "")
}

func createClassField(field entities.JpaField) string {
    paramsMap := map[string]string {
        "{%annotations%}":"",
        "{%field_type%}":field.Type,
        "{%field_name%}":field.Name,
    }
    return utils.FormatString(paramsMap, java.JavaEntityField)
}

func createMapper(class entities.JpaEntity, paramsMap map[string]string) entities.BaseJavaClass {
    mapper := entities.BaseJavaClass{
        Packages : paramsMap["{%mapper_package%}"],
        Imports : utils.FormatString(paramsMap, java.JavaMapper.Imports),
        Annotations : utils.FormatString(paramsMap, java.JavaMapper.Annotations),
        ClassType : java.JavaMapper.ClassType,
        ClassName : class.Name,
        ClassSuffix : paramsMap["{%mapper_suffix%}"],
        Implements : java.JavaMapper.Implements,
        Extends : utils.FormatString(paramsMap, java.JavaMapper.Extends),
    }
    mapper.Body = createMapperBody(class, utils.CopyMap[string, string](paramsMap))    
    mapper.Directory = findDirectoryPath(mapper)
    mapper.FileName = mapper.ClassName + mapper.ClassSuffix + ".java"  
    return mapper
} 

func createMapperBody(object entities.JpaEntity, paramsMap map[string]string) string{
    var dtoSets = []string{}
    var entitySets = []string{}
    for _, field := range object.Fields{
        mapField := map[string]string { "{%field_title%}" : utils.ToTitle(field.Name) }
        rawType := utils.DenestObject(field.Type)
        _, exists := EntityTypes[rawType]
        if !exists {
            dtoSets = append(dtoSets, utils.FormatString(mapField, java.MapperSetDto))
            entitySets = append(entitySets, utils.FormatString(mapField, java.MapperSetEntity))
        }
    }
    paramsMap["{%sets_dto%}"] = strings.Join(dtoSets, "")
    paramsMap["{%sets_entity%}"] = strings.Join(entitySets, "")
    return utils.FormatString(paramsMap, java.JavaMapper.Body)
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

func writeClassImports(classes []entities.JpaEntity){
    for _, class := range classes{
        EntityTypes[class.Name] = createPackage(class, config.CONFIG.DtoPackage)
    }
}

func CreateJavaClasses(){
    classes, err := daos.LoadEntityJson();
    utils.HandleBasicError(err, "Il n'y a pas de fichier de configuration d'entités JPA dans le dossier jpa. Veuillez générer les fichier de configuration à l'aide de la commande jpa avant d'utiliser la commande spring")
    writeClassImports(classes)
    for _, class := range classes{
        paramsMap := createParamsMap(class)
        classList := []entities.BaseJavaClass{
             createController(class, paramsMap),
             createService(class, paramsMap),
             createRepository(class, paramsMap),
             createEntity(class, paramsMap),
             createDto(class, paramsMap),
             createMapper(class, paramsMap),
        }
        writeClassesForOneEntity(classList)
    }
}
