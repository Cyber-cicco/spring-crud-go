package javanalyser

type JavaInterpreted struct {
	JavaPackage PackageStatement
	JavaImports []ImportStatement
	JavaClass   Class
}

type Method struct {
	Final       bool
	Static      bool
    Default     bool
	Abstract    bool
	Visibility  Keyword
	Annotations []Annotation
	ReturnType  JavaType
	Name        Name
	Parameters  []Variable
	Body        Bloc
}

type Bloc struct {
	Instructions    []Instruction
	SubBlocks       []Bloc
	ReturnStatement Variable
}

type Annotation struct {
	Name      Name
	Variables []Variable
}

type Keyword struct {
	Name SyntaxToken
}

type Class struct {
	Name        Name
	Abstract    bool
	Final       bool
	Extends     JavaType
	Implements  []JavaType
	ClassType   Keyword
	Attributes  []Attribute
	Visibility  Keyword
	Methods     []Method
	Classes     []Class
	Annotations []Annotation
}

type Instruction struct {
	Content []SyntaxToken
	Kind    string
}

type Variable struct {
    Annotations []Annotation
	Name     SyntaxToken
	JavaType JavaType
	Value    string
}

type Attribute struct {
	Annotations []Annotation
	Visibility  Keyword
	Final       bool
	Static      bool
	JavaType    JavaType
	Name        Name
	Value       string
	Null        bool
}

type JavaType struct {
	Name     SyntaxToken
	SubTypes []JavaType
}

type ImportStatement struct {
	Keyword     Keyword
	JavaImport  SyntaxToken
	PackagePath []SyntaxToken
}

type PackageStatement struct {
	Keyword     Keyword
	PackagePath []SyntaxToken
}

type Name struct {
	Name SyntaxToken
}

type AttributeI struct {
	Name     string
	JavaType JavaType
}
