package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "spring-go",
  Short: "Génération de code pour Spring et Angular",
  Long: 
`
Génération de code pour Spring et Angular
Permet de créer des classes Java de base.
Possède un utilitaire de création d'entités JPA et de services, DTO et mappers associés.
Possède un utilitaire permettant de créer des interfaces typescript à partir des dtos d'un projet Spring.
Possède un utilitaire permettant de créer des services angular d'appels au back à partir des controllers d'un projet Spring.
`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(jpaCmd)
    rootCmd.AddCommand(classCmd)
    rootCmd.AddCommand(projectCmd)
    rootCmd.AddCommand(ngCmd)
}
