package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaEntity = entities.BaseJavaClass{
    Packages :
``,
    Imports :
`
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Column;
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
`,
}
