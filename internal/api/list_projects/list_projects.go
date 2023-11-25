package list_projects

import (
	"github.com/aakash-rajur/fold-se/internal/store"
	"github.com/aakash-rajur/fold-se/internal/utils"
)

func listProjects(db store.Database, args Args) ([]*ListProjectsResult, error) {
	offset, limit := int64(0), int64(10)

	lpa := &ListProjectsArgs{
		Limit:  utils.Or(args.Limit, &limit),
		Offset: utils.Or(args.Offset, &offset),
	}

	projects, err := lpa.Query(db)

	if err != nil {
		return nil, err
	}

	return projects, nil
}
