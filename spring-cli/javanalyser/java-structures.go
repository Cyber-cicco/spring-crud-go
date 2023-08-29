package javanalyser

type JavaInterpreted struct {
	JavaPackage PackageStatement
	JavaImports []ImportStatement
	JavaClass   Class
}

type Method struct {
	final      bool
	static     bool
	abstract   bool
	visibility Keyword
    annotations []Annotation
	returnType JavaType
	name       Name
	parameters []Variable
	body       Bloc
}

type Bloc struct {
	instructions    []Instruction
	subBlocks       []Bloc
	returnStatement Variable
}

type Annotation struct {
	name      Name
	variables []Variable
}

type Keyword struct {
	name SyntaxToken
}

type Class struct {
	name        Name
	abstract    bool
	final       bool
	extends     JavaType
	implements  []JavaType
	classType   Keyword
	attributes  []Attribute
	visibility  Keyword
	methods     []Method
	classes     []Class
	annotations []Annotation
}

type Instruction struct {
	content []SyntaxToken
	kind    string
}

type Variable struct {
	name     SyntaxToken
	javaType JavaType
	value    string
}

type Attribute struct {
    annotations []Annotation
	visibility Keyword
	final      bool
	static     bool
	javaType   JavaType
	name       Name
	value      string
	null       bool
}

type JavaType struct {
	Name      SyntaxToken
	SubTypes  []JavaType
}

type ImportStatement struct {
	keyword     Keyword
	javaImport  SyntaxToken
	packagePath []SyntaxToken
}

type PackageStatement struct {
	keyword     Keyword
	packagePath []SyntaxToken
}

type Name struct {
	name SyntaxToken
}

type AttributeI struct {
    Name string
    JavaType JavaType
}
