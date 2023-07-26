package java

import (
	"fr.cybercicco/springgo/spring-cli/entities"
)


var javaService = entities.BaseJavaClass{
	Packages:     
``,
	Imports:     
`import java.util.List;`,
	Annotations:
`@Service
@Validated
@RequiredArgsConstructor`,
	ClassType: 
`class`,
	ClassName: 
``,
	ClassSuffix:
`Service`,
	Implements:
``,
	Extends:
``,
	Body:
`
    private final {class_name}Repository {class_name_lower}Repository;
    private final {class_name}Mapper {class_name_lower}Mapper;

    public List<{class_name}Dto> supprimer({class_name}Dto dto){
        return null;
    };

    public List<{class_name}Dto> changer({class_name}Dto dto){
        return null;
    }

    public List<{class_name}Dto> recuperer(){
        return null;
    }   

    public List<{class_name}Dto> creer({class_name}Dto dto){
        return null;
    }
`,
}

