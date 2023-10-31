package auth

import (
	"net/http"

	domain "github.com/Lucas-Linhar3s/debtManager/domain/auth"
	"github.com/gin-gonic/gin"
)

// SignUp takes in a credentials interface{} parameter and returns a *User and an error.
func SignUp(ctx *gin.Context, opts UserCredentials) (interface{}, error) {
	var (
		c    = &http.Client{}
		repo = domain.NewRepository(c)
	)

	result, err := repo.SignUp(opts)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// SignIn takes in a credentials interface{} and returns a pointer to an interface{} and an error.
func SignIn(ctx *gin.Context, opts UserCredentials) (interface{}, error) {
	var (
		c    = &http.Client{}
		repo = domain.NewRepository(c)
	)

	result, err := repo.SignIn(opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Logount(ctx *gin.Context, bearerKey string) error {
	var (
		c    = &http.Client{}
		repo = domain.NewRepository(c)
	)

	if err := repo.SignOut(bearerKey); err != nil {
		return err
	}

	return nil
}
