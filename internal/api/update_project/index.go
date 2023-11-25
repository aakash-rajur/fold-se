package update_project

import (
	"context"
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"time"
)

func UpdateProject(ctx context.Context, args Args) (*Project, error) {
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

	project, err := updateProject(tx, args)

	if err != nil {
		return nil, err
	}

	users, err := updateUsers(tx, args, project, createdAt)

	if err != nil {
		return nil, err
	}

	hashtags, err := updateHashtags(tx, args, project, createdAt)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	p := &Project{
		Project:  project,
		Users:    users,
		Hashtags: hashtags,
	}

	return p, nil
}
