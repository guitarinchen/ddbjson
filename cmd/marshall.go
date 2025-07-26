// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package cmd

import (
	"github.com/guitarinchen/ddbjson/internal/marshall"
	"github.com/guitarinchen/ddbjson/internal/util"
	"github.com/spf13/cobra"
)

var marshallCmd = &cobra.Command{
	Use:   "marshall",
	Short: "Convert normal JSON to DynamoDB JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := util.NewInputReader()
		input, err := reader.Read(args[0])
		if err != nil {
			return err
		}
		return marshall.Marshall(input)
	},
}

func init() {
	rootCmd.AddCommand(marshallCmd)
}
