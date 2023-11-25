package update_project

import (
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/models"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
)

func updateProject(db store.Database, args Args) (*models.Project, error) {
	project, err := store.Find(db, &models.Project{Id: args.Id})

	if err != nil {
		msg := fmt.Sprintf("project with id '%d' not found", *args.Id)

		return nil, errorx.IllegalArgument.Wrap(err, msg)
	}

	project.Name = utils.Or(project.Name, args.Name)

	project.Description = utils.Or(project.Description, args.Description)

	slug := utils.CreateSlug(*project.Name)

	project.Slug = &slug

	err = store.Update(db, project)

	if err != nil {
		return nil, err
	}

	return project, err
}
