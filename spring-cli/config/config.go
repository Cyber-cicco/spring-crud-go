package config

var CONFIG Config  

const RELATIVE_PATH string = "../../"
const JAVA_PATH string = "src/main/java/"

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
}

type PackageOption struct{
    Package string `json:"package"`
    PackagePolicy string`json:"package-policy"`
    Suffix string `json:"suffix"`
}
