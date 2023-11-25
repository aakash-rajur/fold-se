package create_project

import (
	"github.com/aakash-rajur/fold-se/internal/models"
)

type Args struct {
	Name        *string  `json:"name,omitempty" validate:"alpha,required"`
	Description *string  `json:"description,omitempty"`
	Users       []string `json:"users,omitempty" validate:"alpha,required,min=1"`
	Hashtags    []string `json:"hashtags,omitempty" validate:"alphanum,required,min=1"`
}

type Project struct {
	*models.Project
	Users    []*models.User    `json:"users"`
	Hashtags []*models.Hashtag `json:"hashtags"`
}
