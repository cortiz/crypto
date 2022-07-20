/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var filesToHash []string
// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.AddCommand(hashCmd)
  hashCmd.PersistentFlags().StringP("encoder", "e","base64",  "Encoding of the Output avaiable encoders are: base64 , hex, base32, base64url")
	hashCmd.PersistentFlags().StringSliceVarP(&filesToHash,"file", "f", nil, "List of Files to hash")
}
