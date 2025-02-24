package cmd

import (
	"github.com/govindkailas/vault-backup/internal/app"
	"github.com/govindkailas/vault-backup/internal/pkg/s3"
	"github.com/govindkailas/vault-backup/internal/pkg/vault"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup vault secrets using raft snapshot",
	RunE: func(cmd *cobra.Command, args []string) error { // Use RunE for error handling

		vaultCfg := &vault.Config{
			Token:     viper.GetString("vault_token"),
			Address:   viper.GetString("vault_address"),
			Namespace: viper.GetString("vault_namespace"),
			Timeout:   viper.GetDuration("vault_timeout"),
			CACert:    viper.GetString("vault_ca_cert"), // Get CA Cert from viper
		}

		s3Cfg := &s3.Client{
			AccessKey:       viper.GetString("s3_access_key"),
			SecretAccessKey: viper.GetString("s3_secret_key"),
			Region:          viper.GetString("s3_region"),
			Bucket:          viper.GetString("s3_bucket"),
			Endpoint:        viper.GetString("s3_endpoint"),
			FileName:        viper.GetString("s3_filename"),
		}

		return app.Backup(vaultCfg, s3Cfg) // No need for caCertPath argument anymore
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
