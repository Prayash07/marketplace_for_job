// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Company struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	NumberOfEmployees int    `json:"numberOfEmployees"`
}

type CompanyObject struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	NumberOfEmployees int    `json:"numberOfEmployees"`
}

type Education struct {
	ID                 string `json:"id"`
	DegreeName         string `json:"degree_name"`
	NoOfYearsForDegree int    `json:"no_of_years_for_degree"`
}

type Experience struct {
	ID        string    `json:"id"`
	NoOfYears int       `json:"no_of_years"`
	Company   *Company  `json:"company"`
	Position  *Position `json:"position"`
}

type JobAnnouncement struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Company     *Company `json:"company"`
}

type JobAnnouncementObject struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Position    []string `json:"position"`
	Company     string   `json:"company"`
}

type JobAnnouncementPositions struct {
	JobAnnouncement *JobAnnouncement `json:"jobAnnouncement"`
	Position        *Position        `json:"position"`
}

type JobsInCompany struct {
	JobAnnouncement *JobAnnouncement `json:"jobAnnouncement"`
	Company         *Company         `json:"company"`
}

type Position struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Experience      int              `json:"experience"`
	JobAnnouncement *JobAnnouncement `json:"job_announcement"`
}

type User struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Address   string       `json:"address"`
	Gender    Gender       `json:"gender"`
	Education []*Education `json:"education"`
}

type UserObject struct {
	Name               string `json:"name"`
	Address            string `json:"address"`
	DegreeName         string `json:"degree_name"`
	NoOfYearsForDegree int    `json:"no_of_years_for_degree"`
}

type UsersAppliedInVacancy struct {
	ID              string           `json:"id"`
	User            *User            `json:"user"`
	JobAnnouncement *JobAnnouncement `json:"jobAnnouncement"`
	Company         *Company         `json:"company"`
	Position        *Position        `json:"position"`
}

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
