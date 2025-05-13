# SecurePasswd Package for Golang

_This is a work in progress. The features are not yet implemented._

This is the official package in Golang. Each features will be implemented in this package and will use for the development of the SecurePasswd API.

## Features

### Organization

Each organization is a unique entity in the SecurePasswd platform. It is used to group users and vaults.

- [ ] Create organization
- [ ] Get organization
- [ ] Update organization
- [ ] Delete organization

### User

In each organization, there are users. Each user is a unique entity in the SecurePasswd platform. It is used to manage the access to the vaults.

- [ ] Create user
- [ ] Get user
- [ ] Update user
- [ ] Delete user

`We will use a JWT token to authenticate the user. Each user password will be hashed and stored in the database. We use a key derivation function to generate the password hash. In our case, we will use the argon2 algorithm.`
