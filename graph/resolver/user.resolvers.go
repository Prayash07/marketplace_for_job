package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/Prayash07/practice_project/database/models"
	graph "github.com/Prayash07/practice_project/graph/generated"
	"github.com/Prayash07/practice_project/graph/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Education is the resolver for the education field.
func (r *userResolver) Education(ctx context.Context, obj *model.User) ([]*model.Education, error) {
	results, err := models.Educations(qm.Where("UserId=?", obj.ID)).All(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	var educations []*model.Education
	for _, education := range results {
		educations = append(educations, &model.Education{
			ID:                 string(education.ID),
			NoOfYearsForDegree: education.Years,
			DegreeName:         education.DegreeName,
		})
	}
	return educations, nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
