package cmd

import (
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var ngCmd = &cobra.Command{
	Use:   "ng",
	Short: "Permet de générer des interfaces typescript et des services Angular",
	Long: `
Va créer des interfaces typescript en fonction des DTOs d'un projet Spring Boot
Va créer des services d'appels à l'API utilisant le module HTTP de Angular en fonction des constrollers d'un projet Spring boot
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
        services.WriteAngularInterfaceFile()
        services.WriteAngularServiceFile()
    },
}
