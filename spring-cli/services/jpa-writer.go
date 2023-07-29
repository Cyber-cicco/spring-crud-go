package services

import (
	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/entities"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func CreateJpaEntity(jpaCname *string, jpaFields []string){
    jpaEntity := entities.JpaEntity {}
    jpaEntity.Name = *jpaCname
    jpaEntity.Package = config.CONFIG.BasePackage
    jpaEntity.Fields = []entities.JpaField{}
    hydrateEntityFields(&jpaEntity, jpaFields)
    jpaEntity.FileBytes = daos.Jsonify(jpaEntity)
    jpaEntity.FileName = config.CONFIG.JpaJsonFilePath + *jpaCname + ".json"
    daos.WriteEntityJson(jpaEntity)
}

func hydrateEntityFields(entity *entities.JpaEntity, jpaFields []string){
    for _, val := range jpaFields {
        jpaField := entities.JpaField{
            Name : val,
            Type : utils.InfereTypeByName(val),
            Options : []entities.FieldOption{},
        }
        entity.Fields = append(entity.Fields, jpaField)
    }
}
