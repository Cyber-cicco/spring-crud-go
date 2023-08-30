package javanalyser

import (
	"fmt"
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities/enums"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

func PrintLex(tokens [][]SyntaxToken){
    for _, tokenLine := range tokens{
        for _, token := range tokenLine{
            fmt.Print(string(token.Value) + "-")
        }
        fmt.Print("end ")
        fmt.Println()
    }
}

func PrintImport(javaFile JavaInterpreted){
    for _, javaImport := range javaFile.JavaImports {
        fmt.Printf("%+v\n", javaImport)
    }
}

func PrintFile(javaFile JavaInterpreted){
    PrintImport(javaFile)
    PrintClass(javaFile)
}

func PrintClassAnnotations(javaFile JavaInterpreted){
    for _, javaClassAnnotation := range javaFile.JavaClass.Annotations{
        fmt.Printf("%+v\n", javaClassAnnotation)
    }
}

func PrintClass(javaFile JavaInterpreted){
    fmt.Printf("%+v\n", javaFile.JavaClass)
}

func PrintClassAttributes(javaFile JavaInterpreted){
    fmt.Println("Attributes:")
    for _, javaClassAttribute := range javaFile.JavaClass.Attributes{
        fmt.Printf("%+v\n", javaClassAttribute)
    }
}

func PrintTokens(tokens [][]SyntaxToken){
    for _, tokenLine := range tokens{
        for _, token := range tokenLine{
            fmt.Print(string(token.Value) + "-")
        }
        fmt.Print("end ")
        fmt.Println()
    }
}

func PrintClassMethods(javaFile JavaInterpreted){
    fmt.Println("Methods:")
    fmt.Println(len(javaFile.JavaClass.Methods))
    for _, javaClassMethod := range javaFile.JavaClass.Methods{
        fmt.Printf("%+v\n", javaClassMethod)
    }
}

func GetClassAttributes(javaFile JavaInterpreted) []AttributeI{
    attributes := make([]AttributeI, len(javaFile.JavaClass.Attributes))
    for i, javaClassAttribute := range javaFile.JavaClass.Attributes{
        attributes[i].Name = javaClassAttribute.Name.Name.Value
        attributes[i].JavaType = javaClassAttribute.JavaType
    }
    return attributes
}

func GetClassName(javaFile JavaInterpreted) string{
    return javaFile.JavaClass.Name.Name.Value
}

func GetClassPath(javaFile JavaInterpreted) string {
    classPath := ""
    for _, annotation := range javaFile.JavaClass.Annotations {
        if annotation.Name.Name.Value == "RequestMapping" {
            for _, annotationValue := range annotation.Variables {
                if annotationValue.Name.Value == "value" || annotationValue.Name.Value == "path" || annotationValue.Name.Value == enums.DEFAULT_VAR_NAME {
                    classPath = annotationValue.Value
                }
            }
        }
    }
    return classPath
}

func GetMethodPath(method Method, classPath string) string{
    methodPath := ""
    for _, annotation := range method.Annotations {
        if slices.Contains(CONTROLLER_ANNOATIONS, annotation.Name.Name.Value) {
            for _, annotationValue := range annotation.Variables {
                if annotationValue.Name.Value == "value" || annotationValue.Name.Value == "path" || annotationValue.Name.Value == enums.DEFAULT_VAR_NAME {
                    methodPath = annotationValue.Value
                }
            }
        }
    }
    return classPath + methodPath
}

func FindHttpVerb(method Method) string {
    for _, annotation := range method.Annotations {
        if slices.Contains(CONTROLLER_ANNOATIONS, annotation.Name.Name.Value){
            return strings.ToLower(utils.RemoveSuffix(annotation.Name.Name.Value, "Mapping"))
        }
    }
    return ""
}


