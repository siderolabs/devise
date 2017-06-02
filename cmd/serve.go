// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/autonomy/devise/cli"
	"github.com/spf13/cobra"
)

var (
	datastore    string
	backendPort  string
	uiPort       string
	vaultAddress string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Long:  `Serves the UI and backend services.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Start(&cli.ServeOptions{
			Storage:      datastore,
			BackendPort:  backendPort,
			UIPort:       uiPort,
			VaultAddress: vaultAddress,
		})
	},
}

func init() {
	serveCmd.Flags().StringVar(&backendPort, "backend-port", "50000", "The backend listen port")
	serveCmd.Flags().StringVar(&uiPort, "ui-port", "8080", "The UI listen port")
	serveCmd.Flags().StringVar(&datastore, "datastore", "memory", "The datastore used to persist data")
	serveCmd.Flags().StringVar(&vaultAddress, "vault-address", "http://localhost:8200", "The address for Vault")
	RootCmd.AddCommand(serveCmd)
}
