package entities

import "strings"

type BaseJavaClass struct {
    Packages string
    Imports string
    Annotations string
    ClassType string
    ClassName string
    ClassSuffix string
    Implements string
    Extends string
    Body string
}

type EntityTemplate interface {
    FormatParams(params map[string]string, field *string) 
}

func (j BaseJavaClass) FormatParams(params map[string]string, field *string) {
    for key, value := range params {
        *field = strings.ReplaceAll(*field, key, value)
    }
}

