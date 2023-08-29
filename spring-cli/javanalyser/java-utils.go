package javanalyser

import (
	"fmt"
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
    for _, javaClassAnnotation := range javaFile.JavaClass.annotations{
        fmt.Printf("%+v\n", javaClassAnnotation)
    }
}

func PrintClass(javaFile JavaInterpreted){
    fmt.Printf("%+v\n", javaFile.JavaClass)
}

func PrintClassAttributes(javaFile JavaInterpreted){
    fmt.Println("Attributes:")
    for _, javaClassAttribute := range javaFile.JavaClass.attributes{
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
    fmt.Println(len(javaFile.JavaClass.methods))
    for _, javaClassMethod := range javaFile.JavaClass.methods{
        fmt.Printf("%+v\n", javaClassMethod)
    }
}

func GetClassAttributes(javaFile JavaInterpreted) []AttributeI{
    attributes := make([]AttributeI, len(javaFile.JavaClass.attributes))
    for i, javaClassAttribute := range javaFile.JavaClass.attributes{
        attributes[i].Name = javaClassAttribute.name.name.Value
        attributes[i].JavaType = javaClassAttribute.javaType
    }
    return attributes
}

func GetClassName(javaFile JavaInterpreted) string{
    return javaFile.JavaClass.name.name.Value
}
