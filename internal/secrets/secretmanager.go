package secrets

type SecretManager interface {
	CreateSecret(Secret) error
	GetSecrets() (Secrets, error)
}
type Secrets []Secret
type Secret struct {
	Data  map[string]string
	Name  string
	Scope string
}
