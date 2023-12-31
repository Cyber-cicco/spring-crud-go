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
	Name    string
	Urls    []URL
	Http    []HttpMethod
}

type HttpMethod struct {
	HttpVerb   string
	Url        URL
	Args       []Arg
	ReturnType string
}

type Arg struct {
	Scope string
	Type  string
	Name  string
}

type URL struct {
	VarName      string
	AbsolutePath string
	Path         string
}

func mapJavaService(javaFile javanalyser.JavaInterpreted) AngularService {
	angularService := AngularService{}
	angularService.Name = utils.RemoveSuffix(javaFile.JavaClass.Name.Name.Value, config.CONFIG.ControllerPackage.Suffix) + "HttpService"
	classPath := javanalyser.GetClassPath(javaFile)
	createUrls(&angularService, javaFile.JavaClass.Methods, classPath)
	return angularService
}

func createUrls(angularService *AngularService, methods []javanalyser.Method, classPath string) {
	for _, method := range methods {
		if method.Visibility.Name.Value == "public" {
			for _, annotation := range method.Annotations {
				if isRequestAnnotation(annotation) {
					methodPath := javanalyser.GetMethodPath(method)
					httpVerb := strings.ToLower(utils.RemoveSuffix(annotation.Name.Name.Value, "Mapping"))
					if !urlAlreadyExists(angularService, classPath + methodPath) {
						angularService.Urls = append(angularService.Urls, URL{
							VarName:      utils.CreateUrlVarName(methodPath),
							AbsolutePath: classPath + methodPath,
                            Path:         methodPath,
						})
					}
					createMethodForUrl(angularService, method, classPath+methodPath, httpVerb)
				}
			}
		}
	}
}

func createMethodForUrl(angularService *AngularService, method javanalyser.Method, absPath, httpVerb string) {
	httpMethod := HttpMethod{}
	httpMethod.Args = createArgs(method)
	url, err := findUrlForMethod(absPath, angularService)
	utils.HandleTechnicalError(err, "Erreur dans l'analyse d'un fichier Java")
	httpMethod.Url = url
	httpMethod.HttpVerb = httpVerb
	httpMethod.ReturnType = javanalyser.FindTsType(method.ReturnType, map[string]string{}, "")
	angularService.Http = append(angularService.Http, httpMethod)
}

func findUrlForMethod(absPath string, angularService *AngularService) (URL, error) {
	for _, url := range angularService.Urls {
		if url.AbsolutePath == absPath {
			return url, nil
		}
	}
	return URL{}, fmt.Errorf("url not found : %v\n", absPath)
}

func createArgs(method javanalyser.Method) []Arg {
	args := []Arg{}
	for _, variable := range method.Parameters {
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
		if url.AbsolutePath == methodUrl {
			return true
		}
	}
	return false
}
