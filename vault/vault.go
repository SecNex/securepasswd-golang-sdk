package vault

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type CreatedAt time.Time
type ExpiresAt time.Time

type Secret struct {
	ID          string
	Name        string
	Description *string
	Value       string
	Password    *string
	CreatedAt   CreatedAt
	ExpiresAt   *ExpiresAt
}

type Secrets []Secret

type Vault struct {
	OrganizationId *string
	Secrets        Secrets
}

func NewVault(orgId *string) *Vault {
	return &Vault{
		OrganizationId: orgId,
		Secrets:        make(Secrets, 0),
	}
}

func NewExpiresAt(minutes int) ExpiresAt {
	expiresAt := time.Now().Add(time.Duration(minutes) * time.Minute)
	return ExpiresAt(expiresAt)
}

func (e *ExpiresAt) String() string {
	return time.Time(*e).Format(time.RFC3339)
}

func NewSecret(name string, description *string, value string, password *string, expiresAt *ExpiresAt) Secret {
	if password != nil && len(*password) == 0 {
		log.Printf("Password is empty, setting to nil!")
		password = nil
	}

	createdAt := CreatedAt(time.Now())

	return Secret{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Value:       value,
		Password:    password,
		ExpiresAt:   expiresAt,
		CreatedAt:   createdAt,
	}
}

func (v *Vault) Add(secret Secret) {
	v.Secrets = append(v.Secrets, secret)
	log.Printf("Added secret %s to vault", secret.ID)
}

func (v *Vault) Get(id string) *Secret {
	for _, secret := range v.Secrets {
		if secret.ID == id {
			return &secret
		}
	}

	return nil
}

func (v *Vault) Delete(id string) {
	var newSecrets Secrets
	for _, secret := range v.Secrets {
		if secret.ID == id {
			continue
		}

		newSecrets = append(newSecrets, secret)
	}

	v.Secrets = newSecrets
}

func (v *Vault) List() Secrets {
	for _, secret := range v.Secrets {
		log.Printf("%s (%s) - %s", secret.Name, secret.ID, secret.ExpiresAt.String())
	}

	return v.Secrets
}

func (v *Vault) Clear() {
	v.Secrets = make(Secrets, 0)
}
