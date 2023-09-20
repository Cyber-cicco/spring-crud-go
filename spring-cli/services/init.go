package services

import (

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/templates/java"
)

type textTemplate struct {
    name string
    content string
}

/*
*   CreateBaseProject
*   Créer un fichier de configuration du CLI dans le dossier
*   Vérifiera l'existence d'un projet java par la recherche du répertoire src/main/java,
*   et essaiera de trouver le package principal à partir de cela.
*   Créera également des templates permettant de customiser les fichiers générés
*/
func CreateBaseProject(){
    newConfigFile := config.Config {
        BasePackage : daos.FindBasePackage(),
        EreaseFiles : false,
        EntityPackage : config.PackageOption {
            Package: "entities",
            PackagePolicy: "base",
            Suffix: "",
        },
        DtoPackage : config.PackageOption {
            Package: "dto",
            PackagePolicy: "appended",
            Suffix: "Dto",
        },
        MapperPackage : config.PackageOption {
            Package: "dto",
            PackagePolicy: "appended",
            Suffix: "Transformer",
        },
        ServicePackage  : config.PackageOption {
            Package: "services",
            PackagePolicy: "appended",
            Suffix: "Service",
        }, 
        RepositoryPackage : config.PackageOption {
            Package: "repository",
            PackagePolicy: "appended",
            Suffix: "Repository",
        },
        ControllerPackage : config.PackageOption {
            Package: "controller",
            PackagePolicy: "appended",
            Suffix: "Controller",
        },
        TsInterfaceFolder: "models/",
        TsServiceFolder: "http-services/",
    }
    daos.WriteBaseConfigFile(newConfigFile)
    createJavaTemplateTextFiles()

}

/*
*   createJavaTemplateTextFiles
*   Créer les fichiers de templates java
*/
func createJavaTemplateTextFiles() {
        daos.WriteTemplateFile(java.JavaException, "Exception/")
        daos.WriteTemplateFile(java.JavaAnnotation, "Annotation/")
        daos.WriteTemplateFile(java.JavaInterface, "Interface/")
        daos.WriteTemplateFile(java.JavaEnum, "Enum/")
        daos.WriteTemplateFile(java.JavaRecord, "Record/")
        daos.WriteTemplateFile(java.JavaBaseClass, "BaseClass/")
        daos.WriteTemplateFile(java.JavaController, "Controller/")
        daos.WriteTemplateFile(java.JavaService, "Service/")
        daos.WriteTemplateFile(java.JavaDto, "Dto/")
        daos.WriteTemplateFile(java.JavaMapper, "Mapper/")
        daos.WriteTemplateFile(java.JavaEntity, "Entity/")
        daos.WriteTemplateFile(java.JavaRepository, "Repository/")
}


