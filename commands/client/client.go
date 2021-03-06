/*
* Copyright © 2018. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

// Package dapp is the one containing all the cli commands for dapp operations
package client

import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/dovetail-cli/config"
	"github.com/spf13/cobra"
)

func init() {
	ClientCmd.PersistentFlags().StringP("blockchain", "b", config.CORDA, fmt.Sprintf("Target blockchain to deploy to (%s)", strings.Join(config.Blockchains(), "|")))
	ClientCmd.PersistentFlags().StringP("version", "v", "v1.0.0", "client version")
}

// ClientCmd is the command for client app
var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Commands for Client Apps",
	Long:  `Commands for Client Apps`,
}
