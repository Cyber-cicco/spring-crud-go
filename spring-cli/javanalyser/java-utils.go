package javanalyser

import (
	"fmt"
	"strings"

	"fr.cybercicco/springgo/spring-cli/entities/enums"
	"fr.cybercicco/springgo/spring-cli/utils"
	"golang.org/x/exp/slices"
)

/* PrintImport
*
* Print the imports of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintImport(javaFile JavaInterpreted){
    for _, javaImport := range javaFile.JavaImports {
        fmt.Printf("%+v\n", javaImport)
    }
}

/* PrintFile
*
* Print the imports and the class content of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintFile(javaFile JavaInterpreted){
    PrintImport(javaFile)
    PrintClass(javaFile)
}

/* PrintClassAnnotations
*
* Print the annotations of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintClassAnnotations(javaFile JavaInterpreted){
    for _, javaClassAnnotation := range javaFile.JavaClass.Annotations{
        fmt.Printf("%+v\n", javaClassAnnotation)
    }
}

/* PrintClass
*
* Print the content of a class of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintClass(javaFile JavaInterpreted){
    fmt.Printf("%+v\n", javaFile.JavaClass)
}

/* PrintClassAttributes
*
* Print all of the attributes of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintClassAttributes(javaFile JavaInterpreted){
    fmt.Println("Attributes:")
    for _, javaClassAttribute := range javaFile.JavaClass.Attributes{
        fmt.Printf("%+v\n", javaClassAttribute)
    }
}

/* PrintTokens
*
* Print all of the attributes of a java file after parsing
*
* @param javaFile: the file to print
*/
func PrintTokens(tokens [][]SyntaxToken){
    for _, tokenLine := range tokens{
        for _, token := range tokenLine{
            fmt.Print(string(token.Value) + "-")
        }
        fmt.Print("end ")
        fmt.Println()
    }
}

/* PrintLex
*
* Print the filestructure of a java file after lexing
*
* @param tokens: the tokens to print
*/
func PrintClassMethods(javaFile JavaInterpreted){
    fmt.Println("Methods:")
    fmt.Println(len(javaFile.JavaClass.Methods))
    for _, javaClassMethod := range javaFile.JavaClass.Methods{
        fmt.Printf("%+v\n", javaClassMethod)
    }
}

/* GetClassAttributes 
*  extracts attributes from a JavaInterpreted object and returns them as a slice of AttributeI.
*
*  It returns a slice of AttributeI, which represents the attributes of the Java class.
*
*  The function performs the following tasks:
*    - Initializes a slice of AttributeI with the same length as the attributes in the Java class.
*    - Iterates through the attributes of the Java class and populates the AttributeI objects with the respective data.
*    - Copies the name and JavaType from the JavaInterpreted structure to the AttributeI objects.
*
*  @param javaFile (JavaInterpreted): A JavaInterpreted object representing a Java file's parsed structure.
*  @return ([]AttributeI): A slice of AttributeI containing the attributes extracted from the Java class.
*/
func GetClassAttributes(javaFile JavaInterpreted) []AttributeI{
    attributes := make([]AttributeI, len(javaFile.JavaClass.Attributes))
    for i, javaClassAttribute := range javaFile.JavaClass.Attributes{
        attributes[i].Name = javaClassAttribute.Name.Name.Value
        attributes[i].JavaType = javaClassAttribute.JavaType
    }
    return attributes
}

/*GetClassName 
* extracts the name of the Java class from a JavaInterpreted object.
*
* @param javaFile (JavaInterpreted): A JavaInterpreted object representing a Java file's parsed structure.
* @return (string): The name of the Java class.
*/
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

/* GetMethodPath 
*  extracts the API path from a method from its annotations.
*
*  The function iterates through the annotations of the method and checks for specific annotations
*  (e.g., controller annotations) that may contain the API path information.
*
*  @param method (Method): A Method object representing a Java method.
*  @return (string): The path of the API point.
*/
func GetMethodPath(method Method) string{
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
    return methodPath
}

/* FindHttpVerb 
*  extracts the HTTP verb (e.g., GET, POST, etc.) from the method's annotations.
*
*  The function iterates through the annotations of the method and checks for specific annotations
*  (e.g., controller annotations) that may indicate the HTTP verb.
*
*  @param method (Method): A Method object representing a Java method.
*  @return (string): The HTTP verb associated with the method.
*/
func FindHttpVerb(method Method) string {
    for _, annotation := range method.Annotations {
        if slices.Contains(CONTROLLER_ANNOATIONS, annotation.Name.Name.Value){
            return strings.ToLower(utils.RemoveSuffix(annotation.Name.Name.Value, "Mapping"))
        }
    }
    return ""
}


