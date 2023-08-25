package javanalyser

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
    abstract    bool
    final       bool
    extends     JavaType
    implements  []JavaType
    classType   Keyword
    attributes  []Attribute
    visibility  Visibility
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
