package javanalyser

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities/enums"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

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
    j := 0
    fmt.Println(tokens[iPosition][0])
    if tokens[iPosition][j].kind == enums.ANNOTATION_DELIMITER_KIND {
        annotations, j = createAnnotations(tokens[iPosition])
        class.annotations = annotations
        j++
    }
    if !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[iPosition][j].value)) {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[iPosition][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    if tokens[iPosition][j].value == "public" {
        class.visibility = Visibility{visibility : Keyword{name : tokens[iPosition][j]}}
        j++
    } else {
        class.visibility = Visibility{visibility : Keyword{name : SyntaxToken{kind : enums.WORD_KIND, value : "protected"}}}
    }
    if !slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[iPosition][j].value)) {
        switch string(tokens[iPosition][j].value) {
        case "final":
            class.final = true
            j++
        case "abstract":
            class.abstract = true
            j++
        default:
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[iPosition][j].value), config.ERR_JAVA_PARSING_FAILED)
        }
    }
    if slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[iPosition][j].value)) {
        class.classType = Keyword{name : tokens[iPosition][j]}
        j++
    } else {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[iPosition][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    if tokens[iPosition][j].kind != enums.WORD_KIND {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[iPosition][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    class.name = createName(tokens[iPosition][j])
    j++
    if tokens[iPosition][j].kind == enums.OPEN_BRACKET_KIND {
        return class
    }
    if tokens[iPosition][j].value != "extends" {
        j++
        class.extends, j= createJavaType(tokens[iPosition], j)
        j++
    }
    if tokens[iPosition][j].value != "implements" {
        j++
        for tokens[iPosition][j].kind != enums.OPEN_BRACKET_KIND {
            singleImplement, j := createJavaType(tokens[iPosition], j)
            class.implements = append(class.implements, singleImplement)
            j++
            if tokens[iPosition][j].kind != enums.COMMA_KIND {
                utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[iPosition][j].value), config.ERR_JAVA_PARSING_FAILED)
            }
        }
    }
    return class;
}

func createJavaType(tokens []SyntaxToken, j int) (JavaType, int) {
    return JavaType {}, j
}

func createAnnotations(tokens []SyntaxToken) ([]Annotation, int) {
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
    return annotations, j
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
