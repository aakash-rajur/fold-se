package search_projects

import (
	"context"
	lp "github.com/aakash-rajur/fold-se/internal/api/list_projects"
	"github.com/aakash-rajur/fold-se/internal/es"
)

func SearchProject(ctx context.Context, args Args) ([]lp.ListProjectsResult, error) {
	esc, err := es.GetEsc(ctx)

	if err != nil {
		return nil, err
	}

	return searchProjects(esc, args)
}
