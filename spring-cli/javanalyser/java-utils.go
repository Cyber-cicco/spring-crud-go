package javanalyser

import (
	"fmt"
)

func PrintLex(tokens [][]SyntaxToken){
    for _, tokenLine := range tokens{
        for _, token := range tokenLine{
            fmt.Print(string(token.value) + "-")
        }
        fmt.Print("end ")
        fmt.Println()
    }
}

func PrintImport(javaFile JavaFile){
    for _, javaImport := range javaFile.javaImports {
        fmt.Printf("%+v\n", javaImport)
    }
}

func PrintFile(javaFile JavaFile){
    PrintImport(javaFile)
    PrintClassAnnotations(javaFile)
}

func PrintClassAnnotations(javaFile JavaFile){
    for _, javaClassAnnotation := range javaFile.javaClass.annotations{
        fmt.Printf("%+v\n", javaClassAnnotation)
    }
}
