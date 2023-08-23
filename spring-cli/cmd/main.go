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
            jpaExecute(jpaCmd)
        case "j":
            jpaExecute(jpaCmd)
        case "init":
            initExecute(initCmd)
        case "i":
            initExecute(initCmd)
        case "c":
            classExecute(classCmd)
        case "class":
            classExecute(classCmd)
        case "p":
            projectExecute(projectCmd)
        case "project":
            projectExecute(projectCmd)
        default:
            utils.HandleUsageError(errors.New("bad usage"), config.ERR_BAD_ARGS)
            os.Exit(1)
        }
}

func jpaExecute(jpaCmd *flag.FlagSet){
    jpaCname := jpaCmd.String("c", "", "Name of the jpa class")
    jpaFieldsString := jpaCmd.String("f", "", "Fields of the class")
    jpaClear := jpaCmd.Bool("clear", false, "Allows you to clear all jpa config files")
    jpaCmd.Parse(os.Args[2:])
    jpaFields := strings.Split(*jpaFieldsString, " ")
    if ((*jpaCname == "" || *jpaFieldsString == "") && !*jpaClear) || (*jpaClear && (*jpaCname != "" || *jpaFieldsString != "")) {
        utils.HandleUsageError(errors.New("args error"), config.ERR_JPA_ARGS)
    }
    if *jpaClear {
        services.DeleteJpaFiles()
    } else {
        services.CreateJpaEntity(jpaCname, jpaFields)
    }
}

func initExecute(initCmd *flag.FlagSet){
    pkg := initCmd.String("p", "", "Nom du package de base")
    initCmd.Parse(os.Args[2:])
    services.CreateBaseProject(pkg)
}

func classExecute(classCmd *flag.FlagSet){
    classType := classCmd.String("t", "", "Type de la classe que vous voulez créer")
    cname := classCmd.String("c", "", "Nom de la classe que vous voulez créer")
    classCmd.Parse(os.Args[2:])
    if *cname == "" {
        utils.HandleUsageError(errors.New("args error"), config.ERR_CLASS_ARGS)
    }
    services.CreateJavaClass(*cname, *classType)
}

func projectExecute(projectCmd *flag.FlagSet){
    projectCmd.Parse(os.Args[2:])
    fmt.Println("subcommand 'project'")
    services.CreateJavaClasses()
}
