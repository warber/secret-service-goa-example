package secretapi

import (
	"context"
	"fmt"
	"log"
	secret "secretsvc/gen/secret"
	"secretsvc/internal/secrets"
)

// secret service example implementation.
// The example methods log the requests and return zero values.
type secretsrvc struct {
	secretManager *secrets.MemorySecretManager
	logger        *log.Logger
}

// NewSecret returns the secret service implementation.
func NewSecret(logger *log.Logger) secret.Service {
	return &secretsrvc{&secrets.MemorySecretManager{}, logger}
}

// GetSecret implements getSecret.
func (s *secretsrvc) GetSecret(ctx context.Context) (res *secret.GetSecretsResponseBody, err error) {
	secrets, err := s.secretManager.GetSecrets()
	if err != nil {
		return nil, err
	}

	var secretResponseBodies []*secret.GetSecretResponseBody
	for _, s := range secrets {
		secretResponseBodies = append(secretResponseBodies, &secret.GetSecretResponseBody{
			Keys:  nil,
			Name:  &s.Name,
			Scope: &s.Scope,
		})
	}
	res = &secret.GetSecretsResponseBody{
		Secrets: secretResponseBodies,
	}

	return
}

// CreateSecret implements createSecret.
func (s *secretsrvc) CreateSecret(ctx context.Context, p *secret.CreateSecretRequestBody) (*secret.CreateSecretResponseBody, error) {
	err := s.secretManager.CreateSecret(secrets.Secret{
		Data:  p.Data,
		Name:  *p.Name,
		Scope: *p.Scope,
	})
	if err != nil {
		return nil, secret.MakeSecretAlreadyExists(fmt.Errorf("Secret with name %s already exists", *p.Name))
	}
	res := &secret.CreateSecretResponseBody{
		Data:  p.Data,
		Name:  p.Name,
		Scope: p.Scope,
	}
	return res, nil
}
