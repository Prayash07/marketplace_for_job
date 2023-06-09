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

// JobAnnouncement is the resolver for the jobAnnouncement field.
func (r *companyResolver) JobAnnouncement(ctx context.Context, obj *model.Company) ([]*model.JobAnnouncement, error) {
	jobAnnouncements, _ := models.JobAnnouncements(qm.Where("company_id=?", obj.ID)).All(ctx, r.Db)

	var jobAnnouncementList []*model.JobAnnouncement

	for _, jobAnnouncement := range jobAnnouncements {
		jobgqlAnnouncement := &model.JobAnnouncement{
			ID:          jobAnnouncement.ID,
			Title:       jobAnnouncement.Title,
			Description: jobAnnouncement.Description,
			URL:         jobAnnouncement.URL,
		}
		jobAnnouncementList = append(jobAnnouncementList, jobgqlAnnouncement)
	}

	return jobAnnouncementList, nil
}

// Company returns graph.CompanyResolver implementation.
func (r *Resolver) Company() graph.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
