package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaMapper = entities.BaseJavaClass{

    Packages :
`
`,
    Imports :
`
import {%entity_package%}.{%class_name%}{%entity_suffix%};
import {%dto_package%}.{%class_name%}{%dto_suffix%};    

import org.springframework.stereotype.Component;
`,
    Annotations :
`
@Component
`,
    ClassType :
`class`,
    ClassName :
``,
    ClassSuffix :
``,
    Implements :
``,
    Extends :
``,
    Body :
`
    public {%class_name%}{%dto_suffix%} to{%class_name_lower%}{%dto_suffix%}({%class_name%}{%entity_suffix%} entity){
        {%class_name%}{%dto_suffix%} dto = new {%class_name%}{%dto_suffix%}();
        dto.setId(entity.getId());
{%sets_dto%}
        //TODO : implémenter les méthodes pour les champs complexes
        return dto;
    }      

    public {%class_name%}{%entity_suffix%} to{%class_name_lower%}({%class_name%}{%dto_suffix%} dto){
        {%class_name%}{%entity_suffix%} entity = new {%class_name%}{%entity_suffix%}();
        entity.setId(dto.getId());
{%sets_entity%}
        //TODO : implémenter les méthodes pour les champs complexes
        return entity;
    }      
`,
}

var MapperSetDto string = 
`        dto.set{%field_title%}(entity.get{%field_title%}());
`
var MapperSetEntity string =
`        entity.set{%field_title%}(dto.get{%field_title%}());
`
