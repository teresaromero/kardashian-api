package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"kardashian_api/auth"
	"kardashian_api/config"
	"kardashian_api/utils/http_errors"
	"kardashian_api/utils/response"
	"net/http"
)

func CallbackHandler(c *gin.Context) {

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		response.HttpError(c, http_errors.InternalServerError(err))
		return
	}

	code := c.Query("code")
	token, err := authenticator.Config.Exchange(context.Background(), code)
	if err != nil {
		response.HttpError(c, &http_errors.HttpError{StatusCode: http.StatusUnauthorized, Err: err, Message: "No token found"})
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		response.HttpError(c, http_errors.InternalServerError(errors.New("no id_token field in oauth2 token")))
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: config.Auth0ClientId,
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)
	if err != nil {
		response.HttpError(c, http_errors.InternalServerError(errors.New("failed to verify ID Token")))
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if errClaim := idToken.Claims(&profile); errClaim != nil {
		response.HttpError(c, http_errors.InternalServerError(errors.New("failed to get user info")))
		return
	}

	response.SingleResponse(c, token)
	return
}

func LoginHandler(c *gin.Context) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		response.HttpError(c, http_errors.InternalServerError(err))
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		response.HttpError(c, http_errors.InternalServerError(err))
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, authenticator.Config.AuthCodeURL(state))
}
