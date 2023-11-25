package list_projects

import (
	"context"
	"github.com/aakash-rajur/fold-se/internal/store"
)

func ListProject(ctx context.Context, args Args) ([]*ListProjectsResult, error) {
	db, err := store.GetDb(ctx)

	if err != nil {
		return nil, err
	}

	return listProjects(db, args)
}
