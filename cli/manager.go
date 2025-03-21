// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/ultravioletrs/cocos/manager"
)

const (
	serverURL = "server-url"
	serverCA  = "server-ca"
	clientKey = "client-key"
	clientCrt = "client-crt"
	caUrl     = "ca-url"
	logLevel  = "log-level"
)

var (
	agentCVMServerUrl string
	agentCVMServerCA  string
	agentCVMClientKey string
	agentCVMClientCrt string
	agentCVMCaUrl     string
	agentLogLevel     string
)

func (c *CLI) NewCreateVMCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-vm",
		Short:   "Create a new virtual machine",
		Example: `create-vm`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if err := c.InitializeManagerClient(cmd); err != nil {
				printError(cmd, "Failed to connect to manager: %v ❌ ", c.connectErr)
				return
			}
			defer c.Close()

			createReq, err := loadCerts()
			if err != nil {
				printError(cmd, "Error loading certs: %v ❌ ", err)
				return
			}

			createReq.AgentCvmServerUrl = agentCVMServerUrl
			createReq.AgentLogLevel = agentLogLevel
			createReq.AgentCvmCaUrl = agentCVMCaUrl

			cmd.Println("🔗 Creating a new virtual machine")

			res, err := c.managerClient.CreateVm(cmd.Context(), createReq)
			if err != nil {
				printError(cmd, "Error creating virtual machine: %v ❌ ", err)
				return
			}

			cmd.Println(color.New(color.FgGreen).Sprintf("✅ Virtual machine created successfully with id %s and port %s", res.SvmId, res.ForwardedPort))
		},
	}

	cmd.Flags().StringVar(&agentCVMServerUrl, serverURL, "", "CVM server URL")
	cmd.Flags().StringVar(&agentCVMServerCA, serverCA, "", "CVM server CA")
	cmd.Flags().StringVar(&agentCVMClientKey, clientKey, "", "CVM client key")
	cmd.Flags().StringVar(&agentCVMClientCrt, clientCrt, "", "CVM client crt")
	cmd.Flags().StringVar(&agentCVMCaUrl, agentCVMCaUrl, "", "CVM CA service URL")
	cmd.Flags().StringVar(&agentLogLevel, logLevel, "", "Agent Log level")

	return cmd
}

func (c *CLI) NewRemoveVMCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove-vm",
		Short:   "Remove a virtual machine",
		Example: `remove-vm <svm_id>`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := c.InitializeManagerClient(cmd); err == nil {
				defer c.Close()
			}

			if c.connectErr != nil {
				printError(cmd, "Failed to connect to manager: %v ❌ ", c.connectErr)
				return
			}

			cmd.Println("🔗 Removing virtual machine")

			_, err := c.managerClient.RemoveVm(cmd.Context(), &manager.RemoveReq{SvmId: args[0]})
			if err != nil {
				printError(cmd, "Error removing virtual machine: %v ❌ ", err)
				return
			}

			cmd.Println(color.New(color.FgGreen).Sprintf("✅ Virtual machine removed successfully"))
		},
	}
}

func fileReader(path string) ([]byte, error) {
	if path == "" {
		return nil, nil
	}

	return os.ReadFile(path)
}

func loadCerts() (*manager.CreateReq, error) {
	clientKey, err := fileReader(agentCVMClientKey)
	if err != nil {
		return nil, err
	}

	clientCrt, err := fileReader(agentCVMClientCrt)
	if err != nil {
		return nil, err
	}

	serverCA, err := fileReader(agentCVMServerCA)
	if err != nil {
		return nil, err
	}

	return &manager.CreateReq{
		AgentCvmServerCaCert: serverCA,
		AgentCvmClientKey:    clientKey,
		AgentCvmClientCert:   clientCrt,
	}, nil
}
