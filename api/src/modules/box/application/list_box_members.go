package application

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"gitlab.misakey.dev/misakey/msk-sdk-go/ajwt"
	"gitlab.misakey.dev/misakey/msk-sdk-go/merror"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/entrypoints"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/events"
)

type ListBoxMembersRequest struct {
	BoxID string
}

func (req *ListBoxMembersRequest) BindAndValidate(eCtx echo.Context) error {
	req.BoxID = eCtx.Param("id")
	return v.ValidateStruct(req,
		v.Field(&req.BoxID, v.Required, is.UUIDv4),
	)
}

func (bs *BoxApplication) ListBoxMembers(ctx context.Context, genReq entrypoints.Request) (interface{}, error) {
	req := genReq.(*ListBoxMembersRequest)

	// retrieve accesses to filters boxes to return
	acc := ajwt.GetAccesses(ctx)
	if err := events.MustBeMember(ctx, bs.db, req.BoxID, acc.IdentityID); err != nil {
		return nil, merror.Transform(err).Describe("checking membership")
	}

	membersIDs, err := events.ListBoxMembers(
		ctx,
		bs.db,
		req.BoxID,
	)
	if err != nil {
		return nil, merror.Transform(err).Describe("listing box members")
	}

	members := make([]events.SenderView, len(membersIDs))
	i := 0
	for _, id := range membersIDs {
		identity, err := bs.identities.Get(ctx, id)
		if err != nil {
			return nil, merror.Transform(err).Describe("getting identity")
		}

		members[i] = events.NewSenderView(identity)
		i += 1
	}

	return members, nil
}
