package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/secnex/securepasswd-golang-sdk/helper/random"
	"github.com/secnex/securepasswd-golang-sdk/vault"
)

func main() {
	log.Printf("SecNex SecurePasswd Vault")
	orgId := uuid.New().String()
	v := vault.NewVault(&orgId)

	password := random.Generate(30)
	passwordStr := password.String()

	expiresAt := vault.NewExpiresAt(10)

	secret := vault.NewSecret("test", nil, "test", &passwordStr, &expiresAt)

	v.Add(secret)

	log.Printf("Secret: %s", secret.ID)

	v.List()
}
