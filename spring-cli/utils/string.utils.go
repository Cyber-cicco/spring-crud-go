package utils

import (
	"errors"
	"strings"
	"unicode"
)

func ToTitle(s string) string {
    runes := []rune(s)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}

func RemoveSuffix(s string, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        return s[:len(s)-len(suffix)]
    }
    return s
}

func ToAttributeName(s string) string {
    runes := []rune(s)
    runes[0] = unicode.ToLower(runes[0])
    return string(runes)
}

func FormatString(paramsMap map[string]string, s string) string {
    positionOfLastOpenVar := 0
    positionOfLastCloseVar := 0
    var formattedTemplate = []string{}
    for i := 0; i < len(s); i++ {
        if i > 0 && s[i] == '%' && s[i-1] == '{' {
            positionOfLastOpenVar = i - 1
            formattedTemplate = append(formattedTemplate, s[positionOfLastCloseVar:positionOfLastOpenVar])
        } else if i == len(s)-1 {
            formattedTemplate = append(formattedTemplate, s[positionOfLastCloseVar:])
        } else if s[i] == '%' && s[i+1] == '}' {
            positionOfLastCloseVar = i + 2
            formattedVar, ok := paramsMap[s[positionOfLastOpenVar:positionOfLastCloseVar]]
            if !ok {
                HandleTechnicalError(errors.New("variable not found"), "Erreur de formattage du template : la variable "+
                s[positionOfLastOpenVar:positionOfLastCloseVar]+
                " n'a pas été trouvée. Si vous avez changé le code ou le template,"+
                "assurez-vous de la correspondance entre la map crée dans createParamMap du fichier java-writer.go et le template pour lequel l'erreur a été levé")
            }
            formattedTemplate = append(formattedTemplate, formattedVar)
        }
    }
    return strings.Join(formattedTemplate, "")
}

func DenestObject(object string) string {
    positionLastOpen := 0
    for i := 0; i < len(object); i++ {
        if object[i] == '<' {
            positionLastOpen = i
        } else if object[i] == '>' {
            if positionLastOpen == 0 {
                HandleTechnicalError(errors.New("syntax error"), "Caractère non conforme trouvé dans le type de l'objet")
            }
            return object[positionLastOpen+1 : i]
        }
    }
    if positionLastOpen != 0 {
        HandleTechnicalError(errors.New("syntax error"), "Caractère non conforme trouvé dans le type de l'objet")
    }
    return object
}

func GetClassNameAndPackageFromArgs(jpaCname string) (string, string) {
    classPackages := strings.Split(jpaCname, ".")
    cname := classPackages[len(classPackages)-1]
    pname := ""
    for i := 0; i < len(classPackages)-1; i++ {
        pname += "." + classPackages[i]
    }
    return cname, pname
}

func ToInterfaceFileName(interfaceName string) string {
    return strings.ToLower(InsertInStringOnPattern(interfaceName, "-", isUpper)) + ".ts"
}

func isUpper(char byte) bool {
    return unicode.IsUpper(rune(char))
}

func isUrlSpecialChar(char byte) bool {
    return char == '/' || char == '-'
}

func isUrlRemovedChar(char byte) bool {
    return char == '/' || char == '-' || char == '{' || char == '}'
}

func InsertInStringOnPattern(name string, appendedChars string, condition func(char byte) bool) string {
    newName := ""
    positionOfLastPattern := 0
    for i := 0; i < len(name); i++ {
        if condition(name[i])  {
            if i != 0 {
                newName += name[positionOfLastPattern:i] + appendedChars
                positionOfLastPattern = i
            }
        }
        if i == len(name)-1 {
            newName += name[positionOfLastPattern:]
        }
    }
    return newName
}

func CreateUrlVarName(methodPath string) string {
    return strings.ToUpper("URL_" +  RemoveInStringOnPattern(InsertInStringOnPattern(methodPath, "_", isUrlSpecialChar), isUrlRemovedChar))
}

func CreateMethodNameFromUrl(url string) string {
    newName := ""
    positionOfLastPattern := -1
    for i := 0; i < len(url); i++ {
        if isUrlRemovedChar(url[i])  {
            appendedString := url[positionOfLastPattern+1:i]
            if appendedString != "" {
                appendedString = ToTitle(appendedString)
            }
            newName += appendedString
            positionOfLastPattern = i
        }
        if i == len(url)-1 {
            newName += url[positionOfLastPattern+1:]
        }
    }
    return newName
}

func RemoveInStringOnPattern(name string, condition func(char byte) bool) string {
    newName := ""
    positionOfLastPattern := -1
    for i := 0; i < len(name); i++ {
        if condition(name[i])  {
            newName += name[positionOfLastPattern+1:i] 
            positionOfLastPattern = i
        }
        if i == len(name)-1 {
            newName += name[positionOfLastPattern+1:]
        }
    }
    return newName
}
