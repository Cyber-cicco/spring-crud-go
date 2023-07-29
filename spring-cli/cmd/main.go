package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"fr.cybercicco/springgo/spring-cli/utils"
)


func main(){
    daos.LoadConfig()
    handleArgs()
}

func handleArgs(){

    if len(os.Args) < 2 {
        utils.HandleBasicError(errors.New("list index out of range"), "Erreur : le nombre d'arguments doit être supérieur à 2")
    }
    
    jpaCmd := flag.NewFlagSet("jpa", flag.ExitOnError)
    ngCmd := flag.NewFlagSet("ng", flag.ExitOnError)
    springCmd := flag.NewFlagSet("spring", flag.ExitOnError)
    cmpCmd := flag.NewFlagSet("components", flag.ExitOnError)
    mapCmd := flag.NewFlagSet("mappers", flag.ExitOnError)
    flag.Parse()
    jpaCname := jpaCmd.String("cname", "", "Name of the jpa class")
    jpaFieldsString := jpaCmd.String("f", "", "Fields of the class")

    jpaCmd.Parse(os.Args[2:])

    jpaFields := strings.Split(*jpaFieldsString, " ")

    switch os.Args[1] {
        case "jpa":
            if *jpaCname == "" || *jpaFieldsString == "" {
                utils.HandleBasicError(errors.New("args error"), "Erreur : vous devez préciser le nom de l'entité et de ses fields pour utiliser la commande jpa")
            }
            services.CreateJpaEntity(jpaCname, jpaFields)
        case "ng":
            ngCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'ng'")
        case "spring":
            springCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'spring'")
            services.CreateJavaClasses()
        case "components":
            cmpCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'components'")
        case "mappers":
            mapCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'mappers'")
        default:
            fmt.Println("expected 'jpa', 'ng', 'spring', 'components' or 'mapper' subcommands")
            os.Exit(1)
        }
}
