package application

import (
	"context"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/events"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/ajwt"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/merror"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/request"
)

type ListAccessesRequest struct {
	boxID string
}

func (req *ListAccessesRequest) BindAndValidate(eCtx echo.Context) error {
	req.boxID = eCtx.Param("id")
	return v.ValidateStruct(req,
		v.Field(&req.boxID, v.Required, is.UUIDv4),
	)
}

func (bs *BoxApplication) ListAccesses(ctx context.Context, genReq request.Request) (interface{}, error) {
	req := genReq.(*ListAccessesRequest)

	// retrieve accesses to filters boxes to return
	acc := ajwt.GetAccesses(ctx)
	if acc == nil {
		return nil, merror.Unauthorized()
	}
	if err := events.MustBeAdmin(ctx, bs.DB, req.boxID, acc.IdentityID); err != nil {
		return nil, err
	}

	accessEvents, err := events.FindActiveAccesses(ctx, bs.DB, req.boxID)
	if err != nil {
		return nil, merror.Transform(err).Describe("getting sender accesses")
	}

	sendersMap, err := events.MapSenderIdentities(ctx, accessEvents, bs.Identities)
	if err != nil {
		return nil, merror.Transform(err).Describe("retrieving events senders")
	}
	views := make([]events.View, len(accessEvents))
	for i, e := range accessEvents {
		views[i], err = events.FormatEvent(e, sendersMap)
		if err != nil {
			return views, merror.Transform(err).Describe("computing access view")
		}
	}
	return views, nil
}
