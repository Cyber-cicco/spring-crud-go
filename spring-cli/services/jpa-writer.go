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
	paramsMap := map[string]string{
		"{%class_name%}":        utils.ToAttributeName(cname),
		"{%target_class_name%}": fieldName,
	}
	switch strings.Split(annotation[1], ":")[0] {
	case "mtm":
		{
			return checkManyToMany(paramsMap)
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

func checkManyToMany(paramsMap map[string]string) []string {
	key := createMapKey(paramsMap)
	val, ok := MTM_MAP[key]
	if !ok {
		val = utils.FormatString(paramsMap, MANY_TO_MANY)
		MTM_MAP[key] = val
	}
	return []string{val}
}

/**
* prend le nom du field et du nom de la class pour en trouver une clé en faisant un XOR dessus.
* Comme ça, dans le cas d'un autre many to many désignant la même solution mais dans l'autre classe
* cela permettra de récupérer le many to many déjà créé.
 */
func createMapKey(paramsMap map[string]string) uint64 {
	runesCname := []rune(paramsMap["{%class_name%}"])
	runesTarget := []rune(paramsMap["{%target_class_name%}"])
	var sumCname uint64
	var sumTarget uint64
	for _, val := range runesCname {
		sumCname += uint64(val)
	}
	for _, val := range runesTarget {
		sumTarget += uint64(val)
	}
	return sumCname | sumTarget
}
