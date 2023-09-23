package java

import "fr.cybercicco/springgo/spring-cli/entities"

var JavaController = entities.BaseJavaClass{
    Packages :
`
`,
    Imports : 
`
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;    
`,
SpecialImports : 
`
import {%dto_package%}.{%class_name%}{%dto_suffix%};
import org.springframework.http.ResponseEntity;
import {%service_package%}.{%class_name%}{%service_suffix%};
import org.springframework.web.bind.annotation.*;

import java.util.List;
`,
    Annotations : 
`
@RestController
@RequiredArgsConstructor
@RequestMapping("/{%class_name_lower%}")`,
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
   
    private final {%class_name%}{%service_suffix%} {%class_name_lower%}{%service_suffix%};

`,

}
