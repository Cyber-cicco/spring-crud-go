package javanalyser

type JavaFile struct {
	javaPackage PackageStatement
	javaImports []ImportStatement
	javaClass   Class
}

type Method struct {
	final      bool
	static     bool
	abstract   bool
	visibility Visibility
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
	visibility  Visibility
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

type Visibility struct {
	visibility Keyword
}

type Attribute struct {
	visibility Visibility
	final      bool
	static     bool
	javaType   JavaType
	name       Name
	value      string
	null       bool
}

type JavaType struct {
	name      SyntaxToken
	className SyntaxToken
	subTypes  []JavaType
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
