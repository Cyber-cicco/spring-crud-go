package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaMapper = entities.BaseJavaClass{

    Packages :
`
`,
    Imports :
`
import {entity_package}.{class_name}{entity_suffix};
import {dto_package}.{class_name}{dto_suffix};    
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
    {class_name}{dto_suffix} to{class_name_lower}{dto_suffix}({class_name}{entity_suffix} entity){
        {class_name}{dto_suffix} dto = new {class_name}();
{sets}
        //TODO : implémenter les méthodes pour les champs complexes
        return dto;
    }      

`,
}
