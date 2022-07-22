package secrets

import (
	"fmt"
)

type MemorySecretManager struct {
	memory Secrets
}

func (m *MemorySecretManager) CreateSecret(secret Secret) error {
	for _, s := range m.memory {
		if s.Name == secret.Name {
			return fmt.Errorf("Secret already exists")
		}
	}
	m.memory = append(m.memory, secret)
	return nil
}

func (m *MemorySecretManager) GetSecrets() (Secrets, error) {
	return m.memory, nil
}
