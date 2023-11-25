package update_project

import "github.com/aakash-rajur/fold-se/internal/models"

type Args struct {
	Id          *int64   `json:"id,omitempty" validate:"required"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Users       []string `json:"users,omitempty"`
	Hashtags    []string `json:"hashtags,omitempty"`
}

type Project struct {
	*models.Project
	Users    []*models.User    `json:"users"`
	Hashtags []*models.Hashtag `json:"hashtags"`
}
