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

func CreateBaseProject(pkg *string){
    newConfigFile := config.Config {
        BasePackage : *pkg, 
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
    }
    daos.WriteBaseConfigFile(newConfigFile)
    createJavaTemplateTextFiles()

}

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


