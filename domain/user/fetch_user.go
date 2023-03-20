package user

import (
	"context"
	"database/sql"
	"github.com/Prayash07/practice_project/database/models"
	"github.com/Prayash07/practice_project/graph/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FetchUserById(ctx context.Context, id string, db *sql.DB) (*model.User, error) {
	userData, err := models.Users(
		qm.Where("id = ?", id),
	).One(ctx, db)

	if err != nil {
		return nil, err
	}
	//
	//usersEducation, err2 := models.Educations(
	//	qm.Select(
	//		"Education.ID",
	//		"Education.DegreeName",
	//		"Education.Years",
	//	),
	//	//qm.InnerJoin("User ON Education.UserId = User.ID"),
	//	qm.Where("UserID = ?", id),
	//).All(ctx, db)
	//
	//if err2 != nil {
	//	return nil, err2
	//}
	//
	//var educations []*model.Education
	//for _, education := range usersEducation {
	//	educations = append(educations, &model.Education{
	//		DegreeName:         education.DegreeName,
	//		NoOfYearsForDegree: education.Years,
	//	})
	//}
	////
	user := &model.User{
		ID:      userData.ID,
		Name:    userData.Name,
		Address: userData.Address,
	}
	return user, nil
}
