package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/config"
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"fr.cybercicco/springgo/spring-cli/utils"
)

func main(){

    if len(os.Args) < 2 {
        utils.HandleUsageError(errors.New("list index out of range"),config.ERR_NOT_ENOUGH_ARGS_MAIN)
    }
    
    jpaCmd := flag.NewFlagSet(config.JPA_CONFIG_CREATION_ARG, flag.ExitOnError)
    projectCmd := flag.NewFlagSet("project", flag.ExitOnError)
    initCmd := flag.NewFlagSet("init", flag.ExitOnError)
    classCmd := flag.NewFlagSet("class", flag.ExitOnError)

    flag.Parse()

    if os.Args[1] != "init" {
        daos.LoadConfig()
    }

    switch os.Args[1] {
        case "jpa":
            jpaCname := jpaCmd.String("cname", "", "Name of the jpa class")
            jpaFieldsString := jpaCmd.String("f", "", "Fields of the class")
            jpaCmd.Parse(os.Args[2:])
            jpaFields := strings.Split(*jpaFieldsString, " ")
            if *jpaCname == "" || *jpaFieldsString == "" {
                utils.HandleUsageError(errors.New("args error"), config.ERR_JPA_ARGS)
            }
            services.CreateJpaEntity(jpaCname, jpaFields)
        case "init":
            pkg := initCmd.String("package", "", "Nom du package de base")
            initCmd.Parse(os.Args[2:])
            services.CreateBaseProject(pkg)
            fmt.Println(pkg)
        case "class":
            classType := classCmd.String("t", "", "Type de la classe que vous voulez créer")
            cname := classCmd.String("c", "", "Nom de la classe que vous voulez créer")
            classCmd.Parse(os.Args[2:])
            if *cname == "" {
                utils.HandleUsageError(errors.New("args error"), config.ERR_CLASS_ARGS)
            }
            services.CreateJavaClass(*cname, *classType)
        case "project":
            projectCmd.Parse(os.Args[2:])
            fmt.Println("subcommand 'project'")
            services.CreateJavaClasses()
        default:
            utils.HandleUsageError(errors.New("bas usage"), config.ERR_BAD_ARGS)
            os.Exit(1)
        }
}
