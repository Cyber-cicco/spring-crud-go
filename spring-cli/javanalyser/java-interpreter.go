package javanalyser

import (
	"fmt"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/entities/enums"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

/* JAVA_INTERPRETER 
Interpéteur de code java.
Permet de récupérer une structure de code permettant ensuite de parse les différents attributs et méthodes.
N'a pas pour but de faire de l'analyse sémantique.
Doit servir de base pour créer les classes en typescript, le but n'est pas de générer du bytecode java.
*/

var JAVAFILE JavaFile

func OrganizeTokensByMeaning(tokens [][]SyntaxToken) JavaFile {
    i := 0
    JAVAFILE, i = intializedJavaFile(tokens, i)
    JAVAFILE.javaClass, i = createClass(tokens, i)
    return JAVAFILE
}

func createClass(tokens [][]SyntaxToken, i int) (Class, int) {
    class := Class{}
    annotations := []Annotation{}
    j := 0
    fmt.Println(tokens[i][0])
    if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
        annotations, j = createAnnotations(tokens[i])
        class.annotations = annotations
        j++
    }
    if !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[i][j].value)) {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    if tokens[i][j].value == "public" {
        class.visibility = Visibility{visibility : Keyword{name : tokens[i][j]}}
        j++
    } else {
        class.visibility = Visibility{visibility : Keyword{name : SyntaxToken{kind : enums.WORD_KIND, value : "protected"}}}
    }
    if !slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[i][j].value)) {
        switch string(tokens[i][j].value) {
        case "final":
            class.final = true
            j++
        case "abstract":
            class.abstract = true
            j++
        default:
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
        }
    }
    if slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[i][j].value)) {
        class.classType = Keyword{name : tokens[i][j]}
        j++
    } else {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    if tokens[i][j].kind != enums.WORD_KIND {
        utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
    }
    class.name = createName(tokens[i][j])
    j++
    if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
        i++
        createClassBody(tokens, i, &class)
        return class, i
    }
    if tokens[i][j].value != "extends" {
        j++
        class.extends, j= createJavaType(tokens[i], j)
    }
    if tokens[i][j].value != "implements" {
        j++
        for tokens[i][j].kind != enums.OPEN_BRACKET_KIND {
            singleImplement, j := createJavaType(tokens[i], j)
            class.implements = append(class.implements, singleImplement)
            if tokens[i][j].kind != enums.COMMA_KIND {
                utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
            }
        }
    }
    return class, i;
}

func createClassBody(tokens [][]SyntaxToken, i int, class *Class) {
    j := 0
    for tokens[i][j].kind != enums.CLOSE_BRACKET_KIND {
        annotations := []Annotation{}
        if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
            annotations, j = createAnnotations(tokens[i])
            j++
        }
        if slices.Contains(METHOD_IDENTIFIER_KEYWORDS, string(tokens[i][j].value)) {
            k := j
            for tokens[i][k].kind != enums.END_OF_LINE_KIND || tokens[i][k].kind != enums.OPEN_BRACKET_KIND {
                k++
                if k == len(tokens[i]) - 1 {
                    switch string(tokens[i][k].kind) {
                    case enums.END_OF_LINE_KIND:
                        createAttribute(tokens[i],i, j, annotations)
                    case enums.OPEN_BRACKET_KIND:
                        createMethod(tokens[i],i, j, annotations)
                    default:
                        utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].value), config.ERR_JAVA_PARSING_FAILED)
                    }
                }
            }
        }
    }
}

func createMethod(tokens []SyntaxToken, i, j int, annotations []Annotation) (Method, int) {
    return Method{}, i
}

func createAttribute(tokens []SyntaxToken, i, j int, annotations []Annotation) (Attribute, int) {
    return Attribute{}, i
}

func createJavaType(tokens []SyntaxToken, j int) (JavaType, int) {
    javaType := JavaType{}
    if slices.Contains(JAVA_BASE_TYPES, string(tokens[j].value)) {
        javaType.name = tokens[j]
    } else {
        javaType.className = tokens[j]
    }
    j++
    if tokens[j].kind != enums.OPEN_TYPE_KIND {
        return javaType, j
    }
    for tokens[j].kind != enums.CLOSE_TYPE_KIND {
        newType := JavaType{}
        newType, j = createJavaType(tokens, j)
        javaType.subTypes = append(javaType.subTypes, newType)
        if tokens[j].kind != enums.COMMA_KIND && tokens[j].kind != enums.CLOSE_TYPE_KIND {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].value), config.ERR_JAVA_PARSING_FAILED)
        }
        j++
    }
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

func intializedJavaFile(tokens [][]SyntaxToken, i int) (JavaFile, int) {
    javaFile := JavaFile{}
    for !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[i][0].value)) && tokens[i][0].kind != enums.ANNOTATION_DELIMITER_KIND {
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
    return javaFile, i
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
