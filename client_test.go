package vault_test

import (
	"os"
	"testing"

	. "github.com/microdevs/go-support/vault"
)

func TestNewClient_FailsBecauseVaultAddrNotSet(t *testing.T) {
	// when
	_, err := NewClient()

	// then
	if err != ErrNoAddress {
		t.Errorf("expected err: %v to equals %v", ErrNoAddress, err)
	}
}

func TestNewClient_FailsBecauseVaultTokenNotSet(t *testing.T) {
	// given
	os.Setenv("VAULT_ADDR", "http://localhost:8200")

	// when
	_, err := NewClient()

	// then
	if err != ErrNoToken {
		t.Errorf("expected err: %v to equals %v", ErrNoToken, err)
	}
}

func TestVaultClient_Created(t *testing.T) {
	// given
	os.Setenv("VAULT_ADDR", "http://localhost:8200")
	os.Setenv("VAULT_TOKEN", "vault-token")

	// when
	_, err := NewClient()

	// then
	if err != nil {
		t.Errorf("expected client to be created but error arised: %v", err)
	}
}
