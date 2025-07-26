// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ddbjson",
	Short: "Enable inter-conversion between DynamoDB JSON and normal JSON",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
