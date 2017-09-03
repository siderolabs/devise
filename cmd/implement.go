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
	"github.com/autonomy/devise/pkg"
	"github.com/spf13/cobra"
)

var (
	address    string
	plan       string
	vaultToken string
)

// implementCmd represents the implement command
var implementCmd = &cobra.Command{
	Use:   "implement",
	Short: "Implements the plan",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = devise.Implement(&devise.ImplementOptions{
			Address:    address,
			Plan:       plan,
			VaultToken: vaultToken,
		})

		return
	},
}

func init() {
	implementCmd.Flags().StringVar(&address, "address", "localhost:50000", "Specifies the backend address")
	implementCmd.Flags().StringVar(&plan, "plan", "plan.yaml", "Specifies the plan to use")
	implementCmd.Flags().StringVar(&vaultToken, "vault-token", "", "The token used to authenticate to Vault")
	RootCmd.AddCommand(implementCmd)
}
