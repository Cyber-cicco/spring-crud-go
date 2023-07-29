package java

import (
	"fr.cybercicco/springgo/spring-cli/entities"
)


var JavaService = entities.BaseJavaClass{
	Packages:     
``,
	Imports:     
`import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.validation.annotation.Validated;    
import {%dto_package%}.{%class_name%}{%dto_suffix%};
import {%mapper_package%}.{%class_name%}{%mapper_suffix%};
import {%repository_package%}.{%class_name%}{%repository_suffix%};

import java.util.List;
`,
	Annotations:
`@Service
@Validated
@RequiredArgsConstructor`,
	ClassType: 
`class`,
	ClassName: 
``,
	ClassSuffix:
``,
	Implements:
``,
	Extends:
``,
	Body:
`
    private final {%class_name%}{%repository_suffix%} {%class_name_lower%}{%repository_suffix%};
    private final {%class_name%}{%mapper_suffix%} {%class_name_lower%}{%mapper_suffix%};

    public List<{%class_name%}{%dto_suffix%}> supprimer({%class_name%}{%dto_suffix%} dto){
        return null;
    };

    public List<{%class_name%}{%dto_suffix%}> changer({%class_name%}{%dto_suffix%} dto){
        return null;
    }

    public List<{%class_name%}{%dto_suffix%}> recuperer(){
        return null;
    }   

    public List<{%class_name%}{%dto_suffix%}> creer({%class_name%}{%dto_suffix%} dto){
        return null;
    }
`,
}

