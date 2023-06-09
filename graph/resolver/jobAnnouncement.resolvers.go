package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	graph "github.com/Prayash07/practice_project/graph/generated"
	"github.com/Prayash07/practice_project/graph/model"
)

// Company is the resolver for the company field.
func (r *jobAnnouncementResolver) Company(ctx context.Context, obj *model.JobAnnouncement) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: Company - company"))
}

// JobAnnouncement returns graph.JobAnnouncementResolver implementation.
func (r *Resolver) JobAnnouncement() graph.JobAnnouncementResolver {
	return &jobAnnouncementResolver{r}
}

type jobAnnouncementResolver struct{ *Resolver }
