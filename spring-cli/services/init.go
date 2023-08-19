package services

import (
	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
)

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

}
