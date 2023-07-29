package entities

type DtoOption struct {
    Exists bool `json:"exists"`
}

type FieldOption struct {
    Dto DtoOption `json:"dto"`
    Relation string `json:"relation"`
    Annotations []string
}

type JpaField struct {
    Name string `json:"name"`
    Type string`json:"type"`
    Options FieldOption`json:"options"`
}

type JpaEntity struct{
    Name string `json:"name"`
    Package string `json:"package"`
    Fields []JpaField `json:"fields"`
    FileBytes []byte `json:"-"`
    FileName string `json:"-"`
}

