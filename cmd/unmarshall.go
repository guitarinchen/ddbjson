// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package cmd

import (
	"github.com/guitarinchen/ddbjson/internal/unmarshall"
	"github.com/guitarinchen/ddbjson/internal/util"
	"github.com/spf13/cobra"
)

var unmarshallCmd = &cobra.Command{
	Use:   "unmarshall",
	Short: "Convert DynamoDB JSON to normal JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := util.NewInputReader()
		input, err := reader.Read(args[0])
		if err != nil {
			return err
		}
		return unmarshall.Unmarshall(input)
	},
}

func init() {
	rootCmd.AddCommand(unmarshallCmd)
}
