package cmd

import (
	"github.com/govindkailas/vault-backup/internal/app"
	"github.com/govindkailas/vault-backup/internal/pkg/s3"
	"github.com/govindkailas/vault-backup/internal/pkg/vault"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

var forceRestore bool

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a vault backup from raft snapshot",
	RunE: func(cmd *cobra.Command, args []string) error { // Use RunE for error handling

		vaultCfg := &vault.Config{
			Token:        viper.GetString("vault_token"),
			Address:      viper.GetString("vault_address"),
			Namespace:    viper.GetString("vault_namespace"),
			Timeout:      viper.GetDuration("vault_timeout"),
			ForceRestore: forceRestore,
			CACert:       viper.GetString("vault_ca_cert"), // Get CA Cert from viper
		}

		s3Cfg := &s3.Client{
			AccessKey:       viper.GetString("s3_access_key"),
			SecretAccessKey: viper.GetString("s3_secret_key"),
			Region:          viper.GetString("s3_region"),
			Bucket:          viper.GetString("s3_bucket"),
			Endpoint:        viper.GetString("s3_endpoint"),
			FileName:        viper.GetString("s3_filename"),
		}

		return app.Restore(vaultCfg, s3Cfg) // No need for caCertPath argument anymore
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().BoolVarP(&forceRestore, "force", "f", false, "force restore")
}
