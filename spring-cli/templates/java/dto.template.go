package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaDto = entities.BaseJavaClass{
    Packages :
``,
    Imports :
`
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
`,
    Annotations :
`
@AllArgsConstructor
@NoArgsConstructor
@Data
@Builder`,
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
    private Long id;
{%fields%}
`,
    SpecialImports :``,
}
