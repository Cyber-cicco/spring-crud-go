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
		annotations, i, j = createAnnotations(tokens, i, j, &class)
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
		class.Visibility = Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "protected"}}
	}
	for !slices.Contains(CLASS_NAME_KEYWORDS, string(tokens[i][j].Value)) {
		switch string(tokens[i][j].Value) {
		case "final":
			class.Final = true
			j++
		case "abstract":
			class.Abstract = true
			j++
		case "static":
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
		i =createClassBody(tokens, i, &class)
		return class, i
	}
	if tokens[i][j].Value != "extends" {
		j++
		class.Extends, j = createJavaType(tokens[i], j)
	}
	if tokens[i][j].Value == "implements" {
		j++
		for tokens[i][j].kind != enums.OPEN_BRACKET_KIND {
			singleImplement, j := createJavaType(tokens[i], j)
			class.Implements = append(class.Implements, singleImplement)
			if tokens[i][j].kind != enums.COMMA_KIND {
				utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
			}
		}
	}
    i++
    i = createClassBody(tokens, i, &class)
	return class, i
}

func createClassBody(tokens [][]SyntaxToken, i int, class *Class) int {
	j := 0
	for tokens[i][j].kind != enums.CLOSE_BRACKET_KIND {
		if tokens[i][j].kind == enums.COMMENTARY_KIND {
			i++
		}
		annotations := []Annotation{}
		if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
			annotations, i, j = createAnnotations(tokens, i, j, class)
		}
		if slices.Contains(METHOD_IDENTIFIER_KEYWORDS, string(tokens[i][j].Value)) {
			k := j
			l := i
			for tokens[l][k].kind != enums.END_OF_LINE_KIND && tokens[l][k].kind != enums.OPEN_BRACKET_KIND {
				k++
                if tokens[l][k].Value == "class" {
                    newClass := Class{}
                    newClass, i = createClass(tokens, i)
                    class.Classes = append(class.Classes, newClass)
                    j = 0
                    break
                }
				if k == len(tokens[l])-1 {
					switch string(tokens[l][k].kind) {
					case enums.END_OF_LINE_KIND:
						attribute := Attribute{}
						attribute, i = createAttribute(tokens, i, j, annotations)
						class.Attributes = append(class.Attributes, attribute)
						j = 0
						break
					case enums.OPEN_BRACKET_KIND:
						method := Method{}
						method, i = createMethod(tokens, i, j, annotations, class)
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
    return i
}

func createMethod(tokens [][]SyntaxToken, i, j int, annotations []Annotation, class *Class) (Method, int) {
	method := Method{}
	method.Annotations = annotations
    if string(tokens[i][j].Value) == "default" {
        method.Default = true
        j++
    }
	switch string(tokens[i][j].Value) {
        case "public", "private", "protected":
            method.Visibility = Keyword{Name: tokens[i][j]}
            j++
        default:
            if class.ClassType.Name.Value == "interface" {
                method.Visibility = Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "public"}}
            } else {
                method.Visibility = Keyword{Name: SyntaxToken{kind: enums.WORD_KIND, Value: "protected"}}
            }
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
    if tokens[i][j].Value == class.Name.Name.Value && tokens[i][j+1].kind == enums.OPEN_PARENTHESIS_KIND {
        method.Name = createName(tokens[i][j])
    } else {
        method.ReturnType, j = createJavaType(tokens[i], j)
        method.Name = createName(tokens[i][j])
    }
    j++
	if tokens[i][j].kind == enums.OPEN_PARENTHESIS_KIND {
		for tokens[i][j].kind != enums.CLOSE_PARENTHESIS_KIND {
			j++
			if tokens[i][j].kind == enums.WORD_KIND || tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
				variable := Variable{}
				if tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
					variable.Annotations,i, j = createAnnotations(tokens,i, j, class)
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
    nbOpenBracket := 1
    nbCloseBracket := 0
	j := 0
	for nbOpenBracket != nbCloseBracket {
		j = 0
		for j < len(tokens[i]) {
			if tokens[i][j].kind == enums.CLOSE_BRACKET_KIND {
				nbCloseBracket++
			}
			if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
                nbOpenBracket++
			}
			j++
		}
		i++
	}
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
    if tokens[j].kind == enums.OPEN_ARRAY_KIND  {
        j++
        if tokens[j].kind == enums.CLOSE_ARRAY_KIND {
            j++
            javaType.Name.Value += "[]"
        } else {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[j].Value), config.ERR_JAVA_PARSING_FAILED)
        }
    }
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

func createAnnotations(tokens [][]SyntaxToken, i, j int, class *Class) ([]Annotation, int, int) {
	if len(tokens[i]) < 2 {
		utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][0].Value), config.ERR_JAVA_PARSING_FAILED)
	}
	annotations := []Annotation{}
	for tokens[i][j].kind == enums.ANNOTATION_DELIMITER_KIND {
		j++
		annotation := Annotation{}
		annotation.Name = createName(tokens[i][j])
		j++
		if tokens[i][j].kind == enums.OPEN_PARENTHESIS_KIND {
			annotation.Variables, i, j = createAnnotationVariable(tokens, i, j, class, annotation)
			j++
		}
		annotations = append(annotations, annotation)
	}
	return annotations, i, j
}

func handleAnnotationArray(i int, tokens [][]SyntaxToken, variable *Variable) (int, int) {
    i++
    j := 0
    for tokens[i][j].kind != enums.CLOSE_BRACKET_KIND {
        variable.Name.Values = append(variable.Name.Values, tokens[i][j])
        j++
        if tokens[i][j].kind != enums.CLOSE_BRACKET_KIND && tokens[i][j].kind != enums.COMMA_KIND {
            utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
        }
        if tokens[i][j].kind != enums.CLOSE_BRACKET_KIND {
            j++
        }
    }
    i++
    return i, 0
}

func createAnnotationVariable(tokens [][]SyntaxToken,i, j int, class *Class, annotation Annotation) ([]Variable, int, int) {
	j++
	variables := []Variable{}
	for tokens[i][j].kind != enums.CLOSE_PARENTHESIS_KIND {
		variable := Variable{}
		variable.Name.Value = ""
        if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
            i, j = handleAnnotationArray(i, tokens, &variable)
        }
		if tokens[i][j].kind == enums.WORD_KIND {
			j++
			if tokens[i][j].kind == enums.EQUAL_KIND {
                variable.Name = tokens[i][j-1]
                j++
			} else if tokens[i][j].kind == enums.CLOSE_PARENTHESIS_KIND {
                variable.Name.Value = enums.DEFAULT_VAR_NAME
                variable.Value = tokens[i][j-1].Value
                variables = append(variables, variable)
                return variables, i, j
            }
		}

        if tokens[i][j].kind == enums.OPEN_BRACKET_KIND {
            i, j = handleAnnotationArray(i, tokens, &variable)
        }

		if tokens[i][j].kind == enums.STRING_KIND {
			if variable.Name.Value == "" {
				if len(variables) > 0 {
					utils.HandleTechnicalError(fmt.Errorf("Unexpected token %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
				} else {
					variable.Name.Value = enums.DEFAULT_VAR_NAME
				}
			}
			variable.Value = tokens[i][j].Value
		}
		if j == len(tokens[i])-1 {
			utils.HandleTechnicalError(fmt.Errorf("Unexpected end of line %s", tokens[i][j].Value), config.ERR_JAVA_PARSING_FAILED)
		}
        if j != 0 || tokens[i][j].kind == enums.COMMA_KIND{
            j++
        }
		variables = append(variables, variable)
	}
	return variables, i, j
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
