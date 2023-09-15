package cmd

import (
	"fr.cybercicco/springgo/spring-cli/daos"
	"fr.cybercicco/springgo/spring-cli/services"
	"github.com/spf13/cobra"
)

var packageProject string

var projectCmd = &cobra.Command{
	Use:   "pr",
	Short: "Générer une structure de fichiers spring boot",
	Long: `
Va générer différentes entités, dtos, mappers, controllers et services en fonction 
des fichiers de configuration du dossier jpa à la racine du projet
  `,

	Run: func(cmd *cobra.Command, args []string) {
        daos.LoadConfig()
		services.CreateBaseProject(&_package)
    },
}

func init() {
    projectCmd.Flags().StringVarP(&packageProject, "package", "p", "", "Nom du package de base du projet")
}
