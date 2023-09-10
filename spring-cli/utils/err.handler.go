package utils

import (
	"fmt"
	"os"

	"fr.cybercicco/springgo/spring-cli/config"
)


func HandleTechnicalError(err error, message string){
    if(err != nil){
        fmt.Println(message)
        panic(err)
    }
}

func HandleUsageError(err error, message string){
    if(err != nil){
        fmt.Println(message)
        fmt.Println(config.USAGE_MESSAGE)
        os.Exit(1)
    }
}

func Warning(err error, message string){
    if(err != nil){
        fmt.Println(message)
        fmt.Println(err)
    }
}
