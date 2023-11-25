package create_project

import (
	"errors"
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/models"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"time"
)

func createUsers(
	db store.Database,
	args Args,
	project *models.Project,
	createdAt *time.Time,
) ([]*models.User, error) {
	users := make([]*models.User, 0)

	for _, usernameArg := range args.Users {
		errMsg := fmt.Sprintf("unable to insert user with name '%s'", usernameArg)

		user, err := store.Find(db, &models.User{Name: utils.PointerTo(usernameArg)})

		if errors.Is(err, store.ErrNotFound) {
			u := &models.User{
				Name:      utils.PointerTo(usernameArg),
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

	return users, nil
}
