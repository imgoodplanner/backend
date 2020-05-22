package entrypoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.misakey.dev/misakey/msk-sdk-go/merror"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/application"
)

// AuthFlowHTTP provides function to bind to routes interacting with login flow
type AuthFlowHTTP struct {
	service application.SSOService
}

// NewAuthFlowHTTP is AuthFlowHTTP constructor
func NewAuthFlowHTTP(service application.SSOService) *AuthFlowHTTP {
	return &AuthFlowHTTP{
		service: service,
	}
}

// Handles GET /login - init login flow request
func (af AuthFlowHTTP) LoginInit(ctx echo.Context) error {
	// parse parameters
	loginChallenge := ctx.QueryParam("login_challenge")
	if loginChallenge == "" {
		return merror.BadRequest().From(merror.OriQuery).Detail("login_challenge", merror.DVRequired)
	}
	// init login then redirect
	redirectURL := af.service.LoginInit(ctx.Request().Context(), loginChallenge)
	return ctx.Redirect(http.StatusFound, redirectURL)
}

// Handles POST /login/step - perform authentication request for a login flow
func (af AuthFlowHTTP) LoginStep(ctx echo.Context) error {
	cmd := application.LoginStepCmd{}

	if err := ctx.Bind(&cmd); err != nil {
		return merror.BadRequest().From(merror.OriBody).Describe(err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return merror.Transform(err).From(merror.OriBody)
	}

	redirect, err := af.service.LoginStep(ctx.Request().Context(), cmd)
	if err != nil {
		return merror.Transform(err).From(merror.OriBody).Describe("could not step on login flow")
	}
	return ctx.JSON(http.StatusOK, redirect)
}

// Handles GET /consent - init login flow request
func (af AuthFlowHTTP) ConsentInit(ctx echo.Context) error {
	consentChallenge := ctx.QueryParam("consent_challenge")
	if consentChallenge == "" {
		return merror.BadRequest().From(merror.OriQuery).Detail("consent_challenge", merror.DVRequired)
	}
	// init consent then redirect
	redirectURL := af.service.ConsentInit(ctx.Request().Context(), consentChallenge)
	return ctx.Redirect(http.StatusFound, redirectURL)
}
