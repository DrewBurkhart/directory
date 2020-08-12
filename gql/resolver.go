package gql

import (
	"context"
	"github.com/DrewBurkhart/directory/model"
	"github.com/shpota/skmz/db"
	"github.com/shpota/skmz/gql/gen"
)

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Practitioners(ctx context.Context, domain string) ([]*model.Practitioner, error) {
	return r.DB.GetPractitioners(domain)
}
