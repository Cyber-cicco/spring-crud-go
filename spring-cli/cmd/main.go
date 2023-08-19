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

    if len(os.Args) < 2 {
        utils.HandleUsageError(errors.New("list index out of range"), "Erreur : le nombre d'arguments doit être supérieur à 2")
    }
    
    jpaCmd := flag.NewFlagSet("jpa", flag.ExitOnError)
    ngCmd := flag.NewFlagSet("ng", flag.ExitOnError)
    projectCmd := flag.NewFlagSet("project", flag.ExitOnError)
    cmpCmd := flag.NewFlagSet("components", flag.ExitOnError)
    mapCmd := flag.NewFlagSet("mappers", flag.ExitOnError)
    initCmd := flag.NewFlagSet("init", flag.ExitOnError)

    flag.Parse()


    switch os.Args[1] {
        case "jpa":
            jpaCname := jpaCmd.String("cname", "", "Name of the jpa class")
            jpaFieldsString := jpaCmd.String("f", "", "Fields of the class")
            jpaCmd.Parse(os.Args[2:])
            jpaFields := strings.Split(*jpaFieldsString, " ")
            if *jpaCname == "" || *jpaFieldsString == "" {
                utils.HandleUsageError(errors.New("args error"), "Erreur : vous devez préciser le nom de l'entité et de ses fields pour utiliser la commande jpa")
            }
            daos.LoadConfig()
            services.CreateJpaEntity(jpaCname, jpaFields)
        case "ng":
            ngCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'ng'")
        case "init":
            pkg := initCmd.String("package", "", "Nom du package de base")
            initCmd.Parse(os.Args[2:])
            services.CreateBaseProject(pkg)
            fmt.Println(pkg)
        case "project":
            projectCmd.Parse(os.Args[2:])
            daos.LoadConfig()
            fmt.Println("subcommand 'project'")
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
