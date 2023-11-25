package create_project

import (
	"context"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"time"
)

func CreateProject(ctx context.Context, args Args) (*Project, error) {
	db, err := store.GetDb(ctx)

	if err != nil {
		return nil, err
	}

	tx, err := db.Beginx()

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to start a transaction")
	}

	defer func() {
		_ = tx.Rollback()
	}()

	createdAt := utils.PointerTo(time.Now())

	project, err := createProject(tx, args, createdAt)

	if err != nil {
		return nil, err
	}

	users, err := createUsers(tx, args, project, createdAt)

	if err != nil {
		return nil, err
	}

	hashTags, err := createHashtags(tx, args, project, createdAt)

	if err != nil {
		return nil, err
	}

	p := &Project{
		Project:  project,
		Users:    users,
		Hashtags: hashTags,
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return p, nil
}
