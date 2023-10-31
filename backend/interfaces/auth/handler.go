package auth

import (
	"net/http"

	"github.com/Lucas-Linhar3s/debtManager/application/auth"
	"github.com/Lucas-Linhar3s/debtManager/config"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary SignUp
// @Description SignUp
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body auth.UserCredentials true "credentials"
// @Success 201 {object} auth.Res
// @Router /signup [post]
// @Security ApiKeyAuth
func SignUp(ctx *gin.Context) {
	var req auth.UserCredentials

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := auth.SignUp(ctx, req)
	if err != nil {
		config.ResponseWithError(ctx, http.StatusBadGateway, err)
		return
	}

	config.Response(ctx, http.StatusCreated, result)
}

// SignIn godoc
// @Summary SignIn
// @Description SignIn
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body auth.UserCredentials true "credentials"
// @Success 200 {object} auth.Res
// @Router /signin [post]
// @Security ApiKeyAuth
// @Security BearerAuth
func SignIn(ctx *gin.Context) {
	var req auth.UserCredentials

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := auth.SignIn(ctx, req)
	if err != nil {
		config.ResponseWithError(ctx, http.StatusBadGateway, err)
		return
	}

	config.Response(ctx, http.StatusOK, result)
}

// Success godoc
// @Summary Success
// @Description Success
// @Tags Auth
// @Accept json
// @Produce json
// @Router /success [get]
// @Security ApiKeyAuth
func Success(ctx *gin.Context) {
	config.ResponseWithMessage(ctx, http.StatusOK, "Conta confirmada com sucesso!")
}
