package application

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/domain"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/merror"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/oidc"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/request"
)

type BackupArchiveView struct {
	domain.BackupArchive
}

func (sso *SSOService) ListBackupArchives(ctx context.Context, _ request.Request) (interface{}, error) {
	acc := oidc.GetAccesses(ctx)
	// querier must have an account
	if acc == nil || acc.AccountID.IsZero() {
		return nil, merror.Forbidden()
	}

	archives, err := sso.backupArchiveService.ListBackupArchives(ctx, acc.AccountID.String)
	if err != nil {
		return nil, err
	}

	views := make([]BackupArchiveView, len(archives))
	for i, archive := range archives {
		views[i].BackupArchive = archive
	}

	return views, nil
}

type BackupArchiveDataQuery struct {
	archiveID string
}

func (query *BackupArchiveDataQuery) BindAndValidate(eCtx echo.Context) error {
	query.archiveID = eCtx.Param("id")

	return v.ValidateStruct(query,
		v.Field(&query.archiveID, v.Required, is.UUIDv4.Error("archive id must be uuid v4 ")),
	)
}

func (sso *SSOService) GetBackupArchiveData(ctx context.Context, gen request.Request) (interface{}, error) {
	query := gen.(*BackupArchiveDataQuery)
	acc := oidc.GetAccesses(ctx)
	// querier must have an account
	if acc == nil || acc.AccountID.IsZero() {
		return "", merror.Forbidden()
	}

	archive, err := sso.backupArchiveService.GetBackupArchive(ctx, query.archiveID)
	if err != nil {
		return "", merror.Transform(err).Describe("retrieving archive")
	}

	if acc.AccountID.String != archive.AccountID {
		return "", merror.Forbidden()
	}

	if archive.DeletedAt.Valid || archive.RecoveredAt.Valid {
		return "", merror.Gone()
	}

	return archive.Data.String, nil
}

type BackupArchiveDeleteCmd struct {
	archiveID string
	Reason    string `json:"reason"`
}

func (cmd *BackupArchiveDeleteCmd) BindAndValidate(eCtx echo.Context) error {
	if err := eCtx.Bind(cmd); err != nil {
		return merror.BadRequest().From(merror.OriQuery)
	}
	cmd.archiveID = eCtx.Param("id")

	return v.ValidateStruct(cmd,
		v.Field(&cmd.Reason, v.Required, v.In("recovery", "deletion")),
		v.Field(&cmd.archiveID, v.Required, is.UUIDv4),
	)
}

func (sso *SSOService) DeleteBackupArchive(ctx context.Context, gen request.Request) (interface{}, error) {
	cmd := gen.(*BackupArchiveDeleteCmd)
	acc := oidc.GetAccesses(ctx)
	// querier must have an account
	if acc == nil || acc.AccountID.IsZero() {
		return nil, merror.Forbidden()
	}

	archive, err := sso.backupArchiveService.GetBackupArchiveMetadata(ctx, cmd.archiveID)
	if err != nil {
		return nil, merror.Transform(err).Describe("retrieving archive metadata")
	}

	if acc.AccountID.String != archive.AccountID {
		return nil, merror.Forbidden()
	}
	if archive.DeletedAt.Valid || archive.RecoveredAt.Valid {
		return nil, merror.Gone()
	}

	return nil, sso.backupArchiveService.DeleteBackupArchive(ctx, cmd.archiveID, cmd.Reason)
}
