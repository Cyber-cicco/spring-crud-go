package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaEntity = entities.BaseJavaClass{
    Packages :
``,
    Imports :
`
import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Builder;
import lombok.NoArgsConstructor;    
`,
    Annotations :
`
@NoArgsConstructor
@AllArgsConstructor
@Data
@Entity
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
    @Id()
    private Long id;      
{%fields%}

`,
}

var JavaEntityField string = 
`    {%annotations%}private {%field_type%} {%field_name%};
`
