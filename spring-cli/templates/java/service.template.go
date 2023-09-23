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
`,
    SpecialImports : `
import {%dto_package%}.{%class_name%}{%dto_suffix%};
import {%mapper_package%}.{%class_name%}{%mapper_suffix%};
import {%repository_package%}.{%class_name%}{%repository_suffix%};

import java.util.List;`,
	Annotations:
`@Service
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

`,
}

