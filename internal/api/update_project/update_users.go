package update_project

import (
	"errors"
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/models"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"time"
)

func updateUsers(
	db store.Database,
	args Args,
	project *models.Project,
	createdAt *time.Time,
) ([]*models.User, error) {
	lpua := &ListProjectUsersArgs{ProjectId: args.Id}

	projectUsers, err := lpua.Query(db)

	if err != nil {
		msg := fmt.Sprintf("unable to find users belonging to project '%d'", *args.Id)

		return nil, errorx.InternalError.Wrap(err, msg)
	}

	projectUserMap := make(map[string]*ListProjectUsersResult)

	for _, pu := range projectUsers {
		projectUserMap[*pu.UserName] = pu
	}

	projectUserArgsMap := make(map[string]bool)

	users := make([]*models.User, 0)

	for _, projectUserArg := range args.Users {
		// map to lookup for the next loop
		projectUserArgsMap[projectUserArg] = true

		_, ok := projectUserMap[projectUserArg]

		if ok {
			continue
		}

		errMsg := fmt.Sprintf("unable to insert user with name '%s'", projectUserArg)

		user, err := store.Find(db, &models.User{Name: utils.PointerTo(projectUserArg)})

		if errors.Is(err, store.ErrNotFound) {
			u := &models.User{
				Name:      utils.PointerTo(projectUserArg),
				CreatedAt: createdAt,
			}

			err = store.Insert(db, u)

			if err != nil {
				return nil, errorx.InternalError.Wrap(err, errMsg)
			}

			user = u
		} else if err != nil {
			return nil, errorx.InternalError.Wrap(err, errMsg)
		}

		up := &models.UserProject{
			ProjectId: project.Id,
			UserId:    user.Id,
		}

		err = store.Insert(db, up)

		if err != nil {
			return nil, errorx.InternalError.Wrap(err, "unable to associate user with project")
		}

		users = append(users, user)
	}

	for _, projectUser := range projectUsers {
		_, ok := projectUserArgsMap[*projectUser.UserName]

		if ok {
			user := &models.User{
				Id:        projectUser.UserId,
				CreatedAt: projectUser.UserCreatedAt,
				Name:      projectUser.UserName,
			}

			users = append(users, user)

			continue
		}

		err := store.Delete(
			db,
			&models.UserProject{
				ProjectId: projectUser.ProjectId,
				UserId:    projectUser.UserId,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return users, nil
}
