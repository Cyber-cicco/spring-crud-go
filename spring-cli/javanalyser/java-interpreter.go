package javanalyser

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities/enums"
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
    variables []Variable
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
    value    string
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

var iPosition = 0
var JAVAFILE JavaFile

func OrganizeTokensByMeaning(tokens [][]SyntaxToken) JavaFile {
    JAVAFILE = intializedJavaFile(tokens)
    JAVAFILE.javaClass = createClass(tokens)
    return JAVAFILE
}

func createClass(tokens [][]SyntaxToken) Class {
    class := Class{}
    annotations := []Annotation{}
    fmt.Println(tokens[iPosition][0])
    if tokens[iPosition][0].kind == enums.ANNOTATION_DELIMITER_KIND {
        annotations = createAnnotations(tokens[iPosition])
    }
    class.annotations = annotations
    return class;
}

func createAnnotations(tokens []SyntaxToken) []Annotation {
    if(len(tokens) < 2){
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[0].value), config.ERR_JAVA_PARSING_FAILED)
    }
    j := 0
    annotations := []Annotation{}
    for tokens[j].kind == enums.ANNOTATION_DELIMITER_KIND {
        j++
        annotation := Annotation{}
        annotation.name = createName(tokens[j])
        j++
        if tokens[j].kind == enums.OPEN_PARENTHESIS_KIND {
            annotation.variables, j = createAnnotationVariable(tokens, j)
        }
        annotations = append(annotations, annotation)
    }
    return annotations
}

func createAnnotationVariable(tokens []SyntaxToken, j int) ([]Variable, int) {
    j++
    variables := []Variable{}
    for tokens[j].kind != enums.CLOSE_PARENTHESIS_KIND{
        variable := Variable{}
        variable.name.value = ""
        if tokens[j].kind == enums.WORD_KIND {
            variable.name = tokens[j]
            j++
            if tokens[j].kind != enums.EQUAL_KIND{
                utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
            }
            j++
        }
        if tokens[j].kind == enums.STRING_DELIMITER_KIND {
            if variable.name.value == "" {
                if len(variables) > 0 {
                    utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
                } else {
                    variable.name.value = enums.DEFAULT_VAR_NAME
                }
            }
            j++
            variable.value = ""
            for tokens[j].kind != enums.STRING_DELIMITER_KIND && j < len(tokens){
                variable.value += tokens[j].value
                j++
            }
            if j == len(tokens) - 1 {
                utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
            }
        }
        if j == len(tokens) - 1 {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
        }
        j++
        variables = append(variables, variable)
    }
    return variables, j
}

func createName(token SyntaxToken) Name {
    return Name{ name : token }
}

func intializedJavaFile(tokens [][]SyntaxToken) JavaFile {
    javaFile := JavaFile{}
    for !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[iPosition][0].value)) && tokens[iPosition][0].kind != enums.ANNOTATION_DELIMITER_KIND {
        switch string(tokens[iPosition][0].value) {
        case "package":
            javaPackage := createPackageStatement(tokens[iPosition])
            javaFile.javaPackage = javaPackage
        case "import":
            importStatement := createImportStatement(tokens[iPosition])
            javaFile.javaImports = append(javaFile.javaImports, importStatement)
        }
        iPosition++
    }
    return javaFile
}

func createImportStatement(tokens []SyntaxToken) ImportStatement {
    importStatement := ImportStatement{
        keyword : Keyword{name : tokens[0]},
    }
    j := 0
    for tokens[j].kind != enums.END_OF_LINE_KIND {
        if tokens[j].kind == enums.WORD_KIND || tokens[j].kind == enums.STAR_KIND {
            if tokens[j+1].kind == enums.DOT_KIND {
                importStatement.packagePath = append(importStatement.packagePath, tokens[j])
            } else {
                importStatement.javaImport = tokens[j]
            }
        } else if tokens[j].kind != enums.DOT_KIND {
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
    for tokens[j].kind != enums.END_OF_LINE_KIND {
        if tokens[j].kind == enums.WORD_KIND {
            packageStatement.packagePath = append(packageStatement.packagePath, tokens[j])
        } else if tokens[j].kind != enums.DOT_KIND {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
        }
        j++
    }
    return packageStatement
}

func isKeyword(token SyntaxToken) bool {
    if token.kind == enums.WORD_KIND {
        v := token.value
        return slices.Contains(KEYWORDS, string(v))
    }
    return false
}
