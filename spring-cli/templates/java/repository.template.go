package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaRepository = entities.BaseJavaClass{
    Packages :
``,
    Imports :
`
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
`,
    SpecialImports : `
import {%entity_package%}.{%class_name%};`,
    Annotations :
``,
    ClassType :
`interface`,
    ClassName :
``,
    ClassSuffix :
``,
    Implements :
``,
    Extends :
` extends JpaRepository<{%class_name%}, Long> `,
    Body :
``,
}
