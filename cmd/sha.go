/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cortiz/crypto/pkg/hash"
	"github.com/cortiz/crypto/pkg/formaters"
	"github.com/spf13/cobra"
)

var format string

// shaCmd represents the sha command
var shaCmd = &cobra.Command{
	Use:   "sha",
	Short: "Hash the give file or input text",
	Long:  `Computes the SHA2 256 or 512 (default) Hash for the given arguments or files if used with --files`,
	Run: func(cmd *cobra.Command, args []string) {
		outputType := cmd.Flag("output").Value.String()
		shaType := cmd.Flag("type").Value.String()
		encoder := cmd.Flag("encoder").Value.String()
		digestInputs := make(map[string]interface{}, len(args))
	    for i := 0; i < len(args); i++ {
    	digestInputs[args[i]] = hash.SHA(shaType,encoder,[]byte(args[i]))
    }
    if filesToHash != nil {
    	for i:= 0; i< len(filesToHash) ; i++ {
				digestInputs[filesToHash[i]] = hash.SHAFile(shaType, encoder, filesToHash[i])    		
    	}
  }

		fmt.Println(formaters.ProcessOutput(outputType, digestInputs))
	},
}

func init() {
	hashCmd.AddCommand(shaCmd)

	shaCmd.Flags().StringP("type", "t","512",  "SHA 2 Type either 256 or 512 are valid vaules"  )
}
