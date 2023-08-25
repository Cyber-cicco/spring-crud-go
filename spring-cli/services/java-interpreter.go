package services

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

type Interpreted interface {
	constructor()
}

type JavaFile struct {
    javaPackage PackageStatement
    javaImports []ImportStatement
    javaClass Class
}

type Method struct {
	visibility Visibility
	returnType JavaType
	name       SyntaxToken
	body       Block
}

type Block struct {
	instructions    []Instruction
	subBlocks       []Block
	returnStatement Variable
}

type Annotation struct {
	name Name
}

type Keyword struct {
	name SyntaxToken
}

type Class struct {
	name        Name
	attributes  []Attribute
	methods     []Method
	annotations []Annotation
}

type Instruction struct {
	content []SyntaxToken
    kind string
}

type Variable struct {
	name     SyntaxToken
	javaType JavaType
	value    any
}

type Visibility struct {
	visibility Keyword
}

type Attribute struct {
	visibility Visibility
	javaType   JavaType
	name       Name
}

type JavaType struct {
	name     JavaTypeName
	subTypes []JavaType
}

type JavaTypeName struct {
	name      Keyword
	className SyntaxToken
}

type ImportStatement struct {
    keyword Keyword
    javaImport SyntaxToken
    packagePath []SyntaxToken
}

type PackageStatement struct {
    keyword Keyword
    packagePath []SyntaxToken
}

type Name struct {
	name SyntaxToken
}

var i = 0

func OrganizeTokensByMeaning(tokens [][]SyntaxToken) JavaFile {
    javaFile := intializedJavaFile(tokens)
    for _, javaImport := range javaFile.javaImports {
        fmt.Println(javaImport.javaImport.value)
    }
	return javaFile
}

func intializedJavaFile(tokens [][]SyntaxToken) JavaFile {
    javaFile := JavaFile{}
    for !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[i][0].value)) && tokens[i][0].kind != ANNOTATION_DELIMITER_KIND {
        switch string(tokens[i][0].value) {
            case "package":
                javaPackage := createPackageStatement(tokens[i])
                javaFile.javaPackage = javaPackage
            case "import":
                importStatement := createImportStatement(tokens[i])
                javaFile.javaImports = append(javaFile.javaImports, importStatement)
        }
        i++
    }
    return javaFile
}

func createImportStatement(tokens []SyntaxToken) ImportStatement {
    importStatement := ImportStatement{
        keyword : Keyword{name : tokens[0]},
    }
    j := 0
    for tokens[j].kind != END_OF_LINE_KIND {
        if tokens[j].kind == WORD_KIND || tokens[j].kind == STAR_KIND {
            if tokens[j+1].kind == DOT_KIND {
                importStatement.packagePath = append(importStatement.packagePath, tokens[j])
            } else {
                importStatement.javaImport = tokens[j]
            }
        } else if tokens[j].kind != DOT_KIND {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
        }
        j++
    }
    return importStatement
}

func createPackageStatement(tokens []SyntaxToken) PackageStatement {
    packageStatement := PackageStatement{
        keyword : Keyword{name : tokens[0]},
    }
    j := 1
    for tokens[j].kind != END_OF_LINE_KIND {
        if tokens[j].kind == WORD_KIND {
            packageStatement.packagePath = append(packageStatement.packagePath, tokens[j])
        } else if tokens[j].kind != DOT_KIND {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
        }
        j++
    }
    return packageStatement
}

func isKeyword(token SyntaxToken) bool {
    if token.kind == WORD_KIND{
        v := token.value
        return slices.Contains(KEYWORDS, string(v))
    }
    return false
}
