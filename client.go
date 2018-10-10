package vault

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
)

var (
	ErrNoAddress = errors.New(fmt.Sprintf("empty vault address env. variable [%s]", api.EnvVaultAddress))
	ErrNoToken   = errors.New(fmt.Sprintf("empty vault token env. variable [%s]", api.EnvVaultToken))
)

var _ Client = &vaultClient{}

// Client is an interface for vault client
type Client interface {
	WriteData(path string, data map[string]interface{}) error
	ReadData(path string) (map[string]interface{}, error)
}

func NewClient() (Client, error) {
	apiClient, err := apiClient()
	if err != nil {
		return nil, err
	}

	return &vaultClient{apiClient}, nil
}

func apiClient() (*api.Client, error) {
	addr := os.Getenv(api.EnvVaultAddress)

	if addr == "" {
		return nil, ErrNoAddress
	}

	token := os.Getenv(api.EnvVaultToken)

	if token == "" {
		return nil, ErrNoToken
	}

	// create a vault client
	httpClient := &http.Client{}
	ac, err := api.NewClient(&api.Config{Address: addr, HttpClient: httpClient})

	if err != nil {
		return nil, err
	}

	return ac, nil
}

type vaultClient struct {
	*api.Client
}

func (v *vaultClient) WriteData(path string, data map[string]interface{}) error {
	_, err := v.Logical().Write(path, data)
	if err != nil {
		return err
	}
	return nil
}

func (v *vaultClient) ReadData(path string) (map[string]interface{}, error) {
	secret, err := v.Logical().Read(path)
	if err != nil {
		return nil, err
	}

	return secret.Data, nil
}
