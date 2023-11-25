package create_project

import (
	"errors"
	"github.com/aakash-rajur/fold-se/internal/models"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"time"
)

func createProject(db store.Database, args Args, createdAt *time.Time) (*models.Project, error) {
	_, err := store.Find(db, &models.Project{Name: args.Name})

	if !errors.Is(err, store.ErrNotFound) {
		if err != nil {
			return nil, errorx.InternalError.Wrap(err, "unable to determine if project exists")
		}

		return nil, errorx.InternalError.New("project already exists")
	}

	slug := utils.CreateSlug(*args.Name)

	project := &models.Project{
		Name:        args.Name,
		Description: args.Description,
		Slug:        &slug,
		CreatedAt:   createdAt,
	}

	err = store.Insert(db, project)

	if err != nil {
		return nil, err
	}

	return project, nil
}
