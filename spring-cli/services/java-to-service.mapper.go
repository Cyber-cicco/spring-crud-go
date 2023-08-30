package services

import (
	"fmt"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/javanalyser"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

type AngularService struct {
    Imports string
    Name string
    Urls []URL
    Http []HttpMethod
}

type HttpMethod struct {
    HttpVerb string
    Url URL
    Args []Arg
    ReturnType string
}

type Arg struct {
    Scope string
    Type string
    Name string
}

type URL struct {
    VarName string
    Path string
}

func mapJavaService(javaFile javanalyser.JavaInterpreted) AngularService {
    angularService := AngularService{}
    classPath := javanalyser.GetClassPath(javaFile)
    createUrls(&angularService, javaFile.JavaClass.Methods, classPath)
    fmt.Printf("angularService: %v\n", angularService)
    return angularService
}

func createUrls(angularService *AngularService, methods []javanalyser.Method, classPath string){
    for _, method := range methods {
        if method.Visibility.Name.Value == "public" {
            for _, annotation := range method.Annotations {
                if isRequestAnnotation(annotation) {
                    methodPath := javanalyser.GetMethodPath(method, classPath)
                    httpVerb := strings.ToLower(utils.RemoveSuffix(annotation.Name.Name.Value, "Mapping"))
                    if !urlAlreadyExists(angularService, methodPath) {
                        angularService.Urls = append(angularService.Urls, URL{
                            VarName : utils.CreateUrlVarName(methodPath),
                            Path : methodPath,
                        })
                    }
                    createMethodForUrl(angularService, method, methodPath, httpVerb)
                }
            }
        }
    }
}

func createMethodForUrl(angularService *AngularService, method javanalyser.Method,  methodPath, httpVerb string) {
    httpMethod := HttpMethod{}
    httpMethod.Args = createArgs(method)
    url, err := findUrlForMethod(methodPath, angularService)
    utils.HandleTechnicalError(err, "Erreur dans l'analyse d'un fichier Java")
    httpMethod.Url = url
    angularService.Http = append(angularService.Http, httpMethod)
}

func findUrlForMethod(methodPath string, angularService *AngularService) (URL, error) {
    for _, url := range angularService.Urls {
        if url.Path == methodPath {
            return url, nil
        }
    }
    return URL{}, fmt.Errorf("url not found : %v\n", methodPath)
}

func createArgs(method javanalyser.Method) []Arg{
    args := []Arg{}
    for _, variable := range method.Parameters{
        arg := Arg{}
        arg.Name = variable.Name.Value
        arg.Type = javanalyser.FindTsType(variable.JavaType, map[string]string{}, "")
        if len(variable.Annotations) > 0 {
            arg.Scope = variable.Annotations[0].Name.Name.Value
        } else {
            utils.Warning(fmt.Errorf("Unexpected variable %v\n", variable), config.ERR_JAVA_ANALYSING)
        }
        args = append(args, arg)
    }
    return args
}


func isRequestAnnotation(annotation javanalyser.Annotation) bool {
    return slices.Contains(javanalyser.CONTROLLER_ANNOATIONS, annotation.Name.Name.Value)
}

func urlAlreadyExists(angularService *AngularService, methodUrl string) bool {
    for _, url := range angularService.Urls {
        if url.Path == methodUrl {
            return true
        }
    }
    return false
}
