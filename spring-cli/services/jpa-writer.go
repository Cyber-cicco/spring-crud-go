package services

import (
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateJpaEntity(jpaCname *string, jpaFields []string) {
	jpaEntity := entities.JpaEntity{}
	cname, pname := utils.GetClassNameAndPackageFromArgs(*jpaCname)
	jpaEntity.Name = cname
	jpaEntity.Package = config.CONFIG.BasePackage + pname
	jpaEntity.Fields = []entities.JpaField{}
	hydrateEntityFields(&jpaEntity, jpaFields, cname)
	jpaEntity.FileBytes = daos.Jsonify(jpaEntity)
	jpaEntity.FileName = cname + ".json"
	daos.WriteSimpleFile(config.CONFIG.JpaJsonFilePath, jpaEntity.FileName, jpaEntity.FileBytes)
}

func DeleteJpaFiles() {
	daos.DeleteJpaFiles()
}

func hydrateEntityFields(entity *entities.JpaEntity, jpaFields []string, cname string) {
	for _, val := range jpaFields {
		rawName := strings.Split(val, ":")
		rawName = strings.Split(rawName[0], "@")
		typeName := strings.Split(val, "@")[0]
		fieldName := strings.ReplaceAll(rawName[0], "*", "")
		jpaField := entities.JpaField{
			Name: fieldName,
			Type: utils.InfereTypeByName(typeName),
			Options: entities.FieldOption{
				Annotations: createAnnotations(strings.Split(val, "@"), fieldName, cname),
			},
		}
		entity.Fields = append(entity.Fields, jpaField)
	}
}

func createAnnotations(annotation []string, fieldName string, cname string) []string {
	if len(annotation) < 2 {
		return []string{}
	}
	fieldName = strings.ReplaceAll(fieldName, "List", "")
	fieldName = strings.ReplaceAll(fieldName, "Set", "")
    attributeName := utils.ToAttributeName(cname)
	paramsMap := map[string]string{
		"{%class_name%}": attributeName, 
		"{%target_class_name%}": fieldName,
	}
	switch strings.Split(annotation[1], ":")[0] {
	case "mtm":
		{
			return []string{"mtm"}
		}
	case "mto":
		{
			return []string{utils.FormatString(paramsMap, MANY_TO_ONE)}
		}
	case "otm":
		{
			return []string{utils.FormatString(paramsMap, ONE_TO_MANY)}
		}
	default:
		{
			return []string{}
		}
	}
}
