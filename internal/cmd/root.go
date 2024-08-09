package cmd

import (
	"_5gAKA_UE/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "{-c | --config} path",
	DisableFlagsInUseLine: true,
	Short:                 "A simple simulator of UE in 5GAKA protocol",
	RunE:                  handleYamlConfig,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Set path of config file (default for $HOME/.ue.yaml).")
}

func handleYamlConfig(cmd *cobra.Command, args []string) error {
	path, err := cmd.PersistentFlags().GetString("config")
	if err == nil {
		err = config.ReadYamlConfig(path)
	}
	return err
}
