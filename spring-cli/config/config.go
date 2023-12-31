package config

var CONFIG Config  

const RELATIVE_PATH string = "../../"
const JAVA_PATH string = "src/main/java/"


/*
*   Config
*   Structure de configuration du CLI
*/
type Config struct{
    BaseJavaDir string `json:"-"`
    JpaJsonFilePath string `json:"-"`
    BasePackage string `json:"base-package"`
    EreaseFiles bool `json:"erease-files"`
    EntityPackage PackageOption`json:"entity-package"`
    DtoPackage PackageOption`json:"dto-package"`
    MapperPackage PackageOption`json:"mapper-package"`
    ServicePackage PackageOption`json:"service-package"`
    RepositoryPackage PackageOption`json:"repository-package"`
    ControllerPackage PackageOption`json:"controller-package"`
    ExceptionPackage PackageOption`json:"exception-package"`
    DefaultPackage PackageOption`json:"default-package"`
    TsInterfaceFolder string`json:"ts-interface-folder"`
    TsServiceFolder string`json:"ts-service-folder"`
}


/*
*   PackageOption
*   Structure de configuration d'un package d'un type de classe java spécifique
*/
type PackageOption struct{
    Package string `json:"package"`
    PackagePolicy string`json:"package-policy"`
    Suffix string `json:"suffix"`
}
