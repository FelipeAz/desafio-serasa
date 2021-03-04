package interfaces

import (
	"testing"
)

func TestEncryptString(t *testing.T) {
	// Init
	t.Parallel()
	ch := NewCryptoHandler()
	result := "AAAAAAAAAAAAAAAAAAAAACWnqccsz84cIzKv"
	cpf := "47455415893"

	// Execution
	encryptedCpf, _ := ch.EncryptString(cpf)

	// Validation
	if encryptedCpf != result {
		t.Errorf("valor esperado `%s` recebido `%s`", result, encryptedCpf)
	}
}

func TestDecryptString(t *testing.T) {
	// Init
	t.Parallel()
	ch := NewCryptoHandler()
	result := "47455415893"
	encryptedCPF := "AAAAAAAAAAAAAAAAAAAAACWnqccsz84cIzKv"

	// Execution
	cpf, _ := ch.DecryptString(encryptedCPF)

	// Validation
	if cpf != result {
		t.Errorf("valor esperado `%s`. recebido `%s`", result, cpf)
	}
}
