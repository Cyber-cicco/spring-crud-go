package cmd

import (
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var packagePolicy string
var suffix string
var packageConf string
var typeConf string

var configCmd = &cobra.Command{
	Use:   "cf",
	Short: "Permet de gérer le fichier de configuration via la ligne de commande",
	Long: `
Va mettre à jour le fichier de configuration en fonction des options passées
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
        services.ChangeConfig(&suffix, &_package, &packagePolicy, &classType)
    },
}

func init() {
    configCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Suffixe de la classe")
    configCmd.Flags().StringVarP(&packageConf, "package", "p", "", "Package de la classe")
    configCmd.Flags().StringVarP(&packagePolicy, "package-policy", "o", "", "Politique de package de la classe")
    configCmd.Flags().StringVarP(&typeConf, "class-type", "T", "", "Type de la classe à créer")
    configCmd.MarkFlagsRequiredTogether("suffix", "class-type")
    configCmd.MarkFlagsRequiredTogether("package", "class-type")
    configCmd.MarkFlagsRequiredTogether("package-policy", "class-type")
}
