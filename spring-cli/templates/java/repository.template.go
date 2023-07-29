package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaRepository = entities.BaseJavaClass{
    Packages :
``,
    Imports :
`
import {entity_package}.{class_name};
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
`,
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
` extends JpaRepository<{class_name}, Long>`,
    Body :
``,
}
