/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cortiz/crypto/pkg/formaters"
	"github.com/cortiz/crypto/pkg/hash"
	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "MD5 Hash the give file or input text",
	Long:  `Computes the MD5 Hash for the given arguments or files if used with --files`,
	Run: func(cmd *cobra.Command, args []string) {
		outputType := cmd.Flag("output").Value.String()
		encoder := cmd.Flag("encoder").Value.String()
		digestInputs := make(map[string]interface{}, len(args))
		for i := 0; i < len(args); i++ {
			digestInputs[args[i]] = hash.MD5(encoder, []byte(args[i]))
		}
		if filesToHash != nil {
			for i := 0; i < len(filesToHash); i++ {
				digestInputs[filesToHash[i]] = hash.MD5File(encoder, filesToHash[i])
			}
		}
		fmt.Println(formaters.ProcessOutput(outputType, digestInputs))
	},
}

func init() {
	hashCmd.AddCommand(md5Cmd)
}
