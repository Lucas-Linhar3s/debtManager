package auth

import "github.com/Lucas-Linhar3s/debtManager/infrastructure/services/auth"

type IAuth interface {
	SignUp(credentials interface{}) (*auth.Res, error)
	SignIn(credentials interface{}) (interface{}, error)
	RefreshUser(userToken string, refreshToken string) (*interface{}, error)
	ExchangeCode(opts interface{}) (*interface{}, error)
	SendMagicLink(email string) error
	SignInWithProvider(opts interface{}) (*interface{}, error)
	User(userToken string) (*User, error)
	UpdateUser(userToken string, updateData map[string]interface{}) (*User, error)
	ResetPasswordForEmail(email string) error
	SignOut(userToken string) error
	ConvertForDomain(v auth.User) (User, error)
}
