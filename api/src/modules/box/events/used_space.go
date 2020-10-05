package events

import (
	"context"

	"github.com/go-redis/redis/v7"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/entrypoints"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/files"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/quota"
)

type MetadataForUsedSpaceHandler struct {
	OldEventSize int64
	NewEventSize int64
}

func computeUsedSpace(ctx context.Context, e *Event, exec boil.ContextExecutor, redConn *redis.Client, identities entrypoints.IdentityIntraprocessInterface, _ files.FileStorageRepo, metadata Metadata) error {
	msg := metadata.(*Message)

	return quota.UpdateBoxUsedSpace(ctx, exec, e.BoxID, int64(msg.NewSize), int64(msg.OldSize))
}
