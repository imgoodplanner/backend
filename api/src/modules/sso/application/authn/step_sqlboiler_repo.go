package authn

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gitlab.misakey.dev/misakey/backend/api/src/sdk/merror"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/application/oidc"
	"gitlab.misakey.dev/misakey/backend/api/src/modules/sso/repositories/sqlboiler"
)

type StepSQLBoiler struct {
	db *sql.DB
}

func NewAuthnStepSQLBoiler(db *sql.DB) *StepSQLBoiler {
	return &StepSQLBoiler{
		db: db,
	}
}

func (repo StepSQLBoiler) Create(ctx context.Context, authnStep *Step) error {
	// convert domain to sql model
	sqlAuthnStep := sqlboiler.AuthenticationStep{
		IdentityID: authnStep.IdentityID,
		MethodName: string(authnStep.MethodName),
		Metadata:   authnStep.RawJSONMetadata,
	}
	if !authnStep.CreatedAt.IsZero() {
		sqlAuthnStep.CreatedAt = authnStep.CreatedAt
	}
	if err := sqlAuthnStep.Insert(ctx, repo.db, boil.Infer()); err != nil {
		return err
	}
	// copy data potentially created in SQL layer
	authnStep.ID = sqlAuthnStep.ID
	authnStep.CreatedAt = sqlAuthnStep.CreatedAt
	return nil
}

func (repo StepSQLBoiler) CompleteAt(ctx context.Context, id int, completeTime time.Time) error {
	data := sqlboiler.M{sqlboiler.AuthenticationStepColumns.CompleteAt: null.TimeFrom(completeTime)}
	rowsAff, err := sqlboiler.AuthenticationSteps(sqlboiler.AuthenticationStepWhere.ID.EQ(id)).UpdateAll(ctx, repo.db, data)
	if rowsAff == 0 {
		return merror.NotFound().Describe("no rows affected in persistent layer")
	}
	return err
}

func (repo StepSQLBoiler) Last(
	ctx context.Context,
	identityID string,
	methodName oidc.MethodRef,
) (Step, error) {

	authnStep := Step{}

	mods := []qm.QueryMod{
		sqlboiler.AuthenticationStepWhere.IdentityID.EQ(identityID),
		sqlboiler.AuthenticationStepWhere.MethodName.EQ(string(methodName)),
		qm.OrderBy(sqlboiler.AuthenticationStepColumns.CreatedAt + " DESC"),
	}

	sqlAuthnStep, err := sqlboiler.AuthenticationSteps(mods...).One(ctx, repo.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return authnStep, merror.NotFound().Describe(err.Error()).
				Detail("identity_id", merror.DVNotFound).
				Detail("method_name", merror.DVNotFound)
		}
		return authnStep, err
	}

	// build domain model based on sql data
	authnStep.ID = sqlAuthnStep.ID
	authnStep.IdentityID = sqlAuthnStep.IdentityID
	authnStep.MethodName = oidc.MethodRef(sqlAuthnStep.MethodName)
	authnStep.RawJSONMetadata = sqlAuthnStep.Metadata
	authnStep.CreatedAt = sqlAuthnStep.CreatedAt
	authnStep.CompleteAt = sqlAuthnStep.CompleteAt
	authnStep.Complete = sqlAuthnStep.CompleteAt.Valid
	return authnStep, nil
}

func (repo StepSQLBoiler) DeleteIncomplete(ctx context.Context, identityID string) error {
	mods := []qm.QueryMod{
		sqlboiler.AuthenticationStepWhere.IdentityID.EQ(identityID),
		sqlboiler.AuthenticationStepWhere.CompleteAt.IsNull(),
	}
	// ignore no rows affected since not incomplete step deleted means
	// no incomplete steps anymore in storage: the method did its job
	_, err := sqlboiler.AuthenticationSteps(mods...).DeleteAll(ctx, repo.db)
	return err
}

func (repo StepSQLBoiler) Delete(ctx context.Context, stepID int) error {
	mod := sqlboiler.AuthenticationStepWhere.ID.EQ(stepID)

	rowsAff, err := sqlboiler.AuthenticationSteps(mod).DeleteAll(ctx, repo.db)
	if err != nil {
		return err
	}
	if rowsAff == 0 {
		return merror.NotFound().Describe("deleting authn step").
			Detail("id", merror.DVNotFound)
	}
	return nil
}
