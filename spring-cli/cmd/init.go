package cmd

import (
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var _package string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Créer un fichier de configuration du CLI dans le dossier",
	Long: `
Va créer un fichier de configuration du CLI dans le dossier
Vérifiera l'existence d'un projet java par la recherche du répertoire src/main/java,
et essaiera de trouver le package principal à partir de cela.
  `,

	Run: func(cmd *cobra.Command, args []string) {
        services.CreateBaseProject()
    },
}
