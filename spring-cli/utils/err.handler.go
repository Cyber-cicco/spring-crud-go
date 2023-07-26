package utils

import "fmt"

func HandleBasicError(err error, message string){
    if(err != nil){
        fmt.Println(message)
        panic(err)
    }
}
