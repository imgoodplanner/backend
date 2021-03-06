package quota

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"gitlab.misakey.dev/misakey/backend/api/src/modules/box/repositories/sqlboiler"
)

type Quotum struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	IdentityID string    `json:"identity_id"`
	Value      int64     `json:"value"`
	Origin     string    `json:"origin"`
}

func QuotaToDomain(dbQuotum sqlboiler.StorageQuotum) Quotum {
	return Quotum{
		ID:         dbQuotum.ID,
		CreatedAt:  dbQuotum.CreatedAt,
		IdentityID: dbQuotum.IdentityID,
		Value:      dbQuotum.Value,
		Origin:     dbQuotum.Origin,
	}
}

func (q Quotum) ToSQLBoiler() *sqlboiler.StorageQuotum {
	return &sqlboiler.StorageQuotum{
		ID:         q.ID,
		CreatedAt:  q.CreatedAt,
		IdentityID: q.IdentityID,
		Value:      q.Value,
		Origin:     q.Origin,
	}
}

func List(ctx context.Context, exec boil.ContextExecutor, id string) ([]Quotum, error) {
	dbQuota, err := sqlboiler.StorageQuota(sqlboiler.StorageQuotumWhere.IdentityID.EQ(id)).All(ctx, exec)
	if err != nil {
		return nil, err
	}
	if len(dbQuota) == 0 {
		return []Quotum{}, nil
	}

	quota := make([]Quotum, len(dbQuota))
	for idx, quotum := range dbQuota {
		quota[idx] = QuotaToDomain(*quotum)
	}
	return quota, nil
}

func Create(ctx context.Context, exec boil.ContextExecutor, quotum *Quotum) error {
	return quotum.ToSQLBoiler().Insert(ctx, exec, boil.Infer())
}
