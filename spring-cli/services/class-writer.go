package services

import (

	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/services/java-classes"
	"fr.cybercicco/springgo/spring-cli/templates/java"
	"fr.cybercicco/springgo/spring-cli/utils"
)

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

func writeClassesForOneEntity(classes []entities.BaseJavaClass){
    for _, class := range classes{
        daos.WriteJavaClass(class.Directory, class.FileName, createJavaFileBytes(class))        
    }
}


func CreateJavaClasses(){
    classes, err := daos.LoadEntityJson();
    utils.HandleBasicError(err, "Il n'y a pas de fichier de configuration d'entités JPA dans le dossier jpa. Veuillez générer les fichier de configuration à l'aide de la commande jpa avant d'utiliser la commande spring")
    javaclasses.WriteClassImports(classes)
    for _, class := range classes{
        paramsMap := javaclasses.CreateParamsMap(class)
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

