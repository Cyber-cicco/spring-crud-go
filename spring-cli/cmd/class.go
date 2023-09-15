package cmd

import (
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var classType string
var cname string

var classCmd = &cobra.Command{
	Use:   "cl",
	Short: "Générer des classes Typescript",
	Long: `
  Permet de générer une classe Typescript
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
        if deleteFiles {
            services.DeleteJpaFiles()
            os.Exit(1)
        }
        jpaFields := strings.Split(fields, " ")
		services.CreateJpaEntity(&class, jpaFields)
    },
}

func init() {
    classCmd.Flags().StringVarP(&classType, "type", "t", "", "Type de la classe que vous souhaitez créer")
    classCmd.Flags().StringVarP(&cname, "class", "c", "", "Nom de la classe que vous souhaitez créer")
    classCmd.MarkFlagRequired("class")
}
