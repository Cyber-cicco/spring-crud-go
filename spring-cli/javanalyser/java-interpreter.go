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

var JAVAFILE JavaInterpreted

func OrganizeTokensByMeaning(tokens [][]SyntaxToken) JavaInterpreted {
	i := 0
	JAVAFILE, i = intializedJavaFile(tokens, i)
	JAVAFILE.JavaClass, i = createClass(tokens, i)
	return JAVAFILE
}

func createClass(tokens [][]SyntaxToken, i int) (Class, int) {
	class := Class{}
	annotations := []Annotation{}
	j := 0
	if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
		annotations, j = createAnnotations(tokens[i], j)
		class.Annotations = annotations
		j++
	}
	if !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[i][j].Value)) {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	if tokens[i][j].Value == "public" {
		class.Visibility = Keyword{Name: tokens[i][j]}
		j++
	} else {
		class.Visibility =  Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "protected"}}
	}
	if !slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[i][j].Value)) {
		switch string(tokens[i][j].Value) {
		case "final":
			class.Final = true
			j++
		case "abstract":
			class.Abstract = true
			j++
		default:
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
	}
	if slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[i][j].Value)) {
		class.ClassType = Keyword{Name: tokens[i][j]}
		j++
	} else {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	if tokens[i][j].kind != enums.WORD_KIND {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	class.Name = createName(tokens[i][j])
	j++
	if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
		i++
		createClassBody(tokens, i, &class)
		return class, i
	}
	if tokens[i][j].Value != "extends" {
		j++
		class.Extends, j = createJavaType(tokens[i], j)
	}
	if tokens[i][j].Value != "implements" {
		j++
		for tokens[i][j].kind != enums.OPEN_BRACKET_KIND {
			singleImplement, j := createJavaType(tokens[i], j)
			class.Implements = append(class.Implements, singleImplement)
			if tokens[i][j].kind != enums.COMMA_KIND {
				utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
			}
		}
	}
	return class, i
}

func createClassBody(tokens [][]SyntaxToken, i int, class *Class) {
	j := 0
	for tokens[i][j].kind != enums.CLOSE_BRACKET_KIND {
		annotations := []Annotation{}
		if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
			annotations, j = createAnnotations(tokens[i], j)
		}
		if slices.Contains(METHOD_IDENTIFIER_KEYWORDS, string(tokens[i][j].Value)) {
			k := j
			l := i
			for tokens[l][k].kind != enums.END_OF_LINE_KIND && tokens[l][k].kind != enums.OPEN_BRACKET_KIND {
				k++
				if k == len(tokens[l])-1 {
					switch string(tokens[l][k].kind) {
					case enums.END_OF_LINE_KIND:
						if tokens[l][k-1].kind == enums.CLOSE_PARENTHESIS_KIND {
							utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)

						}
						attribute := Attribute{}
						attribute, i = createAttribute(tokens, i, j, annotations)
						class.Attributes = append(class.Attributes, attribute)
						j = 0
						break
					case enums.OPEN_BRACKET_KIND:
						method := Method{}
						method, i = createMethod(tokens, i, j, annotations)
						class.Methods = append(class.Methods, method)
						j = 0
						break
					default:
						utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
					}
				}
			}
		} else {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
	}
}

func createMethod(tokens [][]SyntaxToken, i, j int, annotations []Annotation) (Method, int) {
	method := Method{}
    method.Annotations = annotations
	switch string(tokens[i][j].Value) {
	case "public", "private", "protected":
		method.Visibility =  Keyword{Name: tokens[i][j]}
		j++
	default:
		method.Visibility =  Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "protected"}}
	}
	if tokens[i][j].Value == "static" {
		method.Static = true
		j++
	}
	if tokens[i][j].Value == "abstract" {
		method.Abstract = true
		if method.Static {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
		j++
	}
	method.ReturnType, j = createJavaType(tokens[i], j)
	method.Name = createName(tokens[i][j])
	j++
	if tokens[i][j].kind == enums.OPEN_PARENTHESIS_KIND {
		for tokens[i][j].kind != enums.CLOSE_PARENTHESIS_KIND {
			j++
			if tokens[i][j].kind == enums.WORD_KIND || tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
				variable := Variable{}
                if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
                    variable.Annotations, j = createAnnotations(tokens[i], j)
                }
				variable.JavaType, j = createJavaType(tokens[i], j)
				variable.Name = tokens[i][j]
				j++
				method.Parameters = append(method.Parameters, variable)
			}
		}
	} else {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
	}
    i++
    method.Body, i = createBloc(tokens, i)
	return method, i
}

func createBloc(tokens [][]SyntaxToken, i int) (Bloc, int) {
    bloc := Bloc{}
    j := 0
    for i < len(tokens) {
        j = 0
        for j < len(tokens[i]) {
            if tokens[i][j].kind == enums.CLOSE_BRACKET_KIND {
                i++
                return bloc, i
            }
            if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
                newBloc := Bloc{}
                i++
                newBloc, i = createBloc(tokens, i)
                bloc.SubBlocks = append(bloc.SubBlocks, newBloc)
            }
            j++
        }
        i++
    }
    utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
    return bloc, i
}

func createAttribute(tokens [][]SyntaxToken, i, j int, annotations []Annotation) (Attribute, int) {
	attribute := Attribute{}
    attribute.Annotations = annotations
	switch string(tokens[i][j].Value) {
	case "public", "private", "protected":
		attribute.Visibility = Keyword{Name: tokens[i][j]}
		j++
	default:
		attribute.Visibility = Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "protected"}}
	}
	if tokens[i][j].Value == "static" {
		attribute.Static = true
		j++
	}
	if tokens[i][j].Value == "final" {
		attribute.Final = true
		j++
	}
	attribute.JavaType, j = createJavaType(tokens[i], j)
	attribute.Name = createName(tokens[i][j])
	j++
	if tokens[i][j].kind == enums.END_OF_LINE_KIND {
		attribute.Null = true
		i++
		return attribute, i
	}
	if tokens[i][j].kind == enums.EQUAL_KIND {
		j++
		attribute.Value = ""
		for tokens[i][j].kind != enums.END_OF_LINE_KIND {
			attribute.Value += tokens[i][j].Value
			j++
		}
	} else {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	i++
	return attribute, i
}

func createJavaType(tokens []SyntaxToken, j int) (JavaType, int) {
	javaType := JavaType{}
    javaType.Name = tokens[j]
	j++
	if tokens[j].kind != enums.OPEN_TYPE_KIND {
		return javaType, j
	}
	j++
	for tokens[j].kind != enums.CLOSE_TYPE_KIND {
		newType := JavaType{}
		newType, j = createJavaType(tokens, j)
		javaType.SubTypes = append(javaType.SubTypes, newType)
		if tokens[j].kind != enums.COMMA_KIND && tokens[j].kind != enums.CLOSE_TYPE_KIND {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
	}
	j++
	return javaType, j
}

func createAnnotations(tokens []SyntaxToken, j int) ([]Annotation, int) {
	if len(tokens) < 2 {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[0].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	annotations := []Annotation{}
	for tokens[j].kind == enums.ANNOTATION_DELIMITER_KIND {
		j++
		annotation := Annotation{}
		annotation.Name = createName(tokens[j])
		j++
		if tokens[j].kind == enums.OPEN_PARENTHESIS_KIND {
			annotation.Variables, j = createAnnotationVariable(tokens, j)
		}
		annotations = append(annotations, annotation)
	}
	return annotations, j
}

func createAnnotationVariable(tokens []SyntaxToken, j int) ([]Variable, int) {
	j++
	variables := []Variable{}
	for tokens[j].kind != enums.CLOSE_PARENTHESIS_KIND {
		variable := Variable{}
		variable.Name.Value = ""
		if tokens[j].kind == enums.WORD_KIND {
			variable.Name = tokens[j]
			j++
			if tokens[j].kind != enums.EQUAL_KIND {
				utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
			}
			j++
		}
		if tokens[j].kind == enums.STRING_DELIMITER_KIND {
			if variable.Name.Value == "" {
				if len(variables) > 0 {
					utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
				} else {
					variable.Name.Value = enums.DEFAULT_VAR_NAME
				}
			}
			j++
			variable.Value = ""
			for tokens[j].kind != enums.STRING_DELIMITER_KIND && j < len(tokens) {
				variable.Value += tokens[j].Value
				j++
			}
			if j == len(tokens)-1 {
				utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
			}
		}
		if j == len(tokens)-1 {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
		j++
		variables = append(variables, variable)
	}
	return variables, j
}

func createName(token SyntaxToken) Name {
	return Name{Name: token}
}

// TODO : régler le bug du package
func intializedJavaFile(tokens [][]SyntaxToken, i int) (JavaInterpreted, int) {
	javaFile := JavaInterpreted{}
	for !slices.Contains(CLASS_IDENTIFIER_KEYWORDS, string(tokens[i][0].Value)) && tokens[i][0].kind != enums.ANNOTATION_DELIMITER_KIND {
		switch string(tokens[i][0].Value) {
		case "package":
			javaPackage := createPackageStatement(tokens[i])
			javaFile.JavaPackage = javaPackage
		case "import":
			importStatement := createImportStatement(tokens[i])
			javaFile.JavaImports = append(javaFile.JavaImports, importStatement)
		}
		i++
	}
	return javaFile, i
}

func createImportStatement(tokens []SyntaxToken) ImportStatement {
	importStatement := ImportStatement{
		Keyword: Keyword{Name: tokens[0]},
	}
	j := 0
	for tokens[j].kind != enums.END_OF_LINE_KIND {
		if tokens[j].kind == enums.WORD_KIND || tokens[j].kind == enums.STAR_KIND {
			if tokens[j+1].kind == enums.DOT_KIND {
				importStatement.PackagePath = append(importStatement.PackagePath, tokens[j])
			} else {
				importStatement.JavaImport = tokens[j]
			}
		} else if tokens[j].kind != enums.DOT_KIND {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
		j++
	}
	return importStatement
}

func createPackageStatement(tokens []SyntaxToken) PackageStatement {
	packageStatement := PackageStatement{
		Keyword: Keyword{Name: tokens[0]},
	}
	j := 1
	for tokens[j].kind != enums.END_OF_LINE_KIND {
		if tokens[j].kind == enums.WORD_KIND {
			packageStatement.PackagePath = append(packageStatement.PackagePath, tokens[j])
		} else if tokens[j].kind != enums.DOT_KIND {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
		j++
	}
	return packageStatement
}

func isKeyword(token SyntaxToken) bool {
	if token.kind == enums.WORD_KIND {
		v := token.Value
		return slices.Contains(KEYWORDS, string(v))
	}
	return false
}
