package auth

import (
	"net/http"

	"github.com/Lucas-Linhar3s/debtManager/infrastructure/services/auth"
)

type repository struct {
	Client *auth.Auth
}

func NewRepository(c *http.Client) IAuth {
	return &repository{
		Client: &auth.Auth{
			Client: c,
		},
	}
}

func (r *repository) ConvertForDomain(v auth.User) (User, error) {
	return User{
		ID:                 v.ID,
		Aud:                v.Aud,
		Email:              v.Email,
		Role:               v.Role,
		InvitedAt:          v.InvitedAt,
		ConfirmedAt:        v.ConfirmedAt,
		ConfirmationSentAt: v.ConfirmationSentAt,
		UserMetadata:       v.UserMetadata,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
	}, nil
}

// SignUp takes in a credentials interface{} parameter and returns a *User and an error.
func (r *repository) SignUp(credentials interface{}) (*auth.Res, error) {
	return r.Client.SignUp(credentials)
}

// SignIn takes in a credentials interface{} and returns a pointer to an interface{} and an error.
func (r *repository) SignIn(credentials interface{}) (interface{}, error) {
	return r.Client.SignIn(credentials)
}

// RefreshUser takes in a userToken and a refreshToken as parameters.
func (r *repository) RefreshUser(userToken string, refreshToken string) (*interface{}, error) {
	return nil, nil
}

// ExchangeCode exchanges a code for a token.
func (r *repository) ExchangeCode(opts interface{}) (*interface{}, error) {
	return nil, nil
}

// SendMagicLink sends a magic link to the specified email address.
func (r *repository) SendMagicLink(email string) error {
	return nil
}

// SignInWithProvider takes in an opts interface{} parameter and returns a *interface{} and an error.
func (r *repository) SignInWithProvider(opts interface{}) (*interface{}, error) {
	return nil, nil
}

// User returns a User and an error.
func (r *repository) User(userToken string) (*User, error) {
	return nil, nil
}

// UpdateUser updates a user in the repository.
func (r *repository) UpdateUser(userToken string, updateData map[string]interface{}) (*User, error) {
	return nil, nil
}

// ResetPasswordForEmail resets the password for the given email.
func (r *repository) ResetPasswordForEmail(email string) error {
	return nil
}

// SignOut takes a userToken string as a parameter and returns an error.
func (r *repository) SignOut(userToken string) error {
	return nil
}
