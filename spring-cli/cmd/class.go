package cmd

import (

	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var classType string
var cname string

var classCmd = &cobra.Command{
	Use:   "cl",
	Short: "Générer des classes Java",
	Long: `
  Permet de générer une classe Java et de préciser son type.
  Peut être une classe simple, une annotation, une interface, un enum, un record ou un repository, un mapper, un service, un controller un DTO.
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
		services.CreateJavaClass(cname, classType)
    },
}

func init() {
    classCmd.Flags().StringVarP(&classType, "type", "t", "", "Type de la classe que vous souhaitez créer")
    classCmd.Flags().StringVarP(&cname, "class", "c", "", "Nom de la classe que vous souhaitez créer")
    classCmd.MarkFlagRequired("class")
}
