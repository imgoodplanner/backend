package application

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/merror"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/request"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/application/authn"
)

// AuthenticationStepCmd orders:
// - the retry of an authentication step init for the identity
type AuthenticationStepCmd struct {
	LoginChallenge string     `json:"login_challenge"`
	Step           authn.Step `json:"authn_step"`
}

func (cmd *AuthenticationStepCmd) BindAndValidate(eCtx echo.Context) error {
	if err := eCtx.Bind(cmd); err != nil {
		return merror.BadRequest().From(merror.OriBody).Describe(err.Error())
	}

	if err := v.ValidateStruct(&cmd.Step,
		v.Field(&cmd.Step.IdentityID, v.Required, is.UUIDv4.Error("identity_id must be an UUIDv4")),
		v.Field(&cmd.Step.MethodName, v.Required),
	); err != nil {
		return err
	}

	return v.ValidateStruct(cmd,
		v.Field(&cmd.LoginChallenge, v.Required),
	)
}

// This method is used to try to init an authentication step
func (sso *SSOService) InitAuthnStep(ctx context.Context, genReq request.Request) (interface{}, error) {
	cmd := genReq.(*AuthenticationStepCmd)

	// 0. check if the identity exists and authable
	identity, err := sso.identityService.Get(ctx, cmd.Step.IdentityID)
	if err != nil {
		return nil, err
	}
	if !identity.IsAuthable {
		return nil, merror.Forbidden().Describe("identity not authable")
	}

	// 1. check login challenge
	_, err = sso.authFlowService.GetLoginContext(ctx, cmd.LoginChallenge)
	if err != nil {
		return nil, merror.NotFound().Describe("finding login challenge").Detail("login_challenge", merror.DVNotFound)
	}

	// 2. we try to init the authentication step
	return nil, sso.AuthenticationService.InitStep(ctx, identity, cmd.Step.MethodName)
}
