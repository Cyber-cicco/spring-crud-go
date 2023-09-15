package cmd

import (
	"os"
	"strings"

	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var class string
var fields string
var deleteFiles bool

var jpaCmd = &cobra.Command{
	Use:   "jpa",
	Short: "Générer et supprimer des fichiers de configuration d'entités JPA",
	Long: `
Permet de générer des fichiers de configuration d'entités JPA dans un dossier jpa à la racine du projet.    
Permet de créer des fichiers servant ensuite à la commande project pour générer toute 
une structure de fichiers dépendante de ces fichiers de configuration
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
        if deleteFiles {
            services.DeleteJpaFiles()
            os.Exit(1)
        }
        if !deleteFiles && (class == "" || fields == "") {
            cmd.Help()
            os.Exit(1)
        }
        jpaFields := strings.Split(fields, " ")
		services.CreateJpaEntity(&class, jpaFields)
    },
}

func init() {
    jpaCmd.Flags().StringVarP(&class, "class", "c", "", "Nom de la classe à générer")
    jpaCmd.Flags().StringVarP(&fields, "fields", "f", "", "Liste des champs de la classe")
    jpaCmd.Flags().BoolVarP(&deleteFiles, "delete", "d", false, "Supprimer les fichiers de configurations JPA")
    jpaCmd.MarkFlagsMutuallyExclusive("class", "delete")
    jpaCmd.MarkFlagsMutuallyExclusive("fields", "delete")
    jpaCmd.MarkFlagsRequiredTogether("class", "fields")
}
