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


// shaCmd represents the sha command
var sha3Cmd = &cobra.Command{
	Use:   "sha3",
	Short: "Hash the give file or input text",
	Long:  `Computes the SHA3 223 256 ,384 or 512 (default) Hash for the given arguments or files if used with --files`,
	Run: func(cmd *cobra.Command, args []string) {
		outputType := cmd.Flag("output").Value.String()
		shaType := cmd.Flag("type").Value.String()
		encoder := cmd.Flag("encoder").Value.String()
		digestInputs := make(map[string]interface{}, len(args))
	    for i := 0; i < len(args); i++ {
    	digestInputs[args[i]] = hash.SHA3(shaType,encoder,[]byte(args[i]))
    }

    if filesToHash != nil {
    	for i:= 0; i< len(filesToHash) ; i++ {
				digestInputs[filesToHash[i]] = hash.SHA3File(shaType, encoder, filesToHash[i])    		
    	}
  }

		fmt.Println(formaters.ProcessOutput(outputType, digestInputs))
	},
}

func init() {
	hashCmd.AddCommand(sha3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	sha3Cmd.Flags().StringP("type", "t","512",  "SHA 3 Type either 224, 256, 384 or 512 are valid values"  )
}
