package services

import (
	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
)

func ChangeConfig(setSuffix, setPackage, setPackagePolicy, classType *string){
    switch *classType{
        case "ent":
            changeOneClassConfig(&config.CONFIG.EntityPackage, setSuffix, setPackage, setPackagePolicy)
        case "dto":
            changeOneClassConfig(&config.CONFIG.DtoPackage, setSuffix, setPackage, setPackagePolicy)
        case "map":
            changeOneClassConfig(&config.CONFIG.MapperPackage, setSuffix, setPackage, setPackagePolicy)
        case "srv": 
            changeOneClassConfig(&config.CONFIG.ServicePackage, setSuffix, setPackage, setPackagePolicy)
        case "repo":
            changeOneClassConfig(&config.CONFIG.RepositoryPackage, setSuffix, setPackage, setPackagePolicy)
        case "ctrl":
            changeOneClassConfig(&config.CONFIG.ControllerPackage, setSuffix, setPackage, setPackagePolicy)
        case "ex":
            changeOneClassConfig(&config.CONFIG.ExceptionPackage, setSuffix, setPackage, setPackagePolicy)
    }
}

func changeOneClassConfig(option *config.PackageOption, setSuffix, setPackage, setPackagePolicy *string){
    if *setSuffix != "" {
        option.Suffix = *setSuffix
    }
    if *setPackage != "" {
        option.Package = *setPackage
    }
    if *setPackagePolicy != "" {
        option.PackagePolicy = *setPackagePolicy
    }
    daos.WriteBaseConfigFile(config.CONFIG)
}
