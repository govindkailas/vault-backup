package vault

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"time"
)

const DEFAULT_VAULT_NAMESPACE = "admin"

type Client struct {
	vaultClient  *vault.Client
	forceRestore bool
}

type Config struct {
	Address      string
	Token        string
	Namespace    string
	ForceRestore bool
	TmpPath      string
	FileName     string
	Timeout      time.Duration
	CACert       string
}

func NewClient(config *Config) (*Client, error) {
	vaultConfig := vault.DefaultConfig()

	vaultConfig.Address = config.Address

	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	if config.CACert != "" { // Check if CACert is provided
		tlsConfig := &vault.TLSConfig{
			CACert: config.CACert,
		}
		err = vaultConfig.ConfigureTLS(tlsConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to configure TLS: %w", err)
		}
	}

	client.SetToken(config.Token)
	client.SetClientTimeout(config.Timeout)

	if config.Namespace == "" {
		config.Namespace = DEFAULT_VAULT_NAMESPACE
	}

	if config.Namespace == "" {
		config.Namespace = DEFAULT_VAULT_NAMESPACE
	}

	client.SetNamespace(config.Namespace)

	return &Client{
		vaultClient:  client,
		forceRestore: config.ForceRestore,
	}, nil
}
