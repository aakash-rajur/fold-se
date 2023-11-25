package search_projects

import (
	"bytes"
	"errors"
	lp "github.com/aakash-rajur/fold-se/internal/api/list_projects"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/joomcode/errorx"
	"net/http"
)

func searchProjects(esc *es.Client, args Args) ([]lp.ListProjectsResult, error) {
	query, err := inferQuery(args)

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to infer query")
	}

	response, err := esc.Search(
		esc.Search.WithIndex("projects"),
		esc.Search.WithBody(bytes.NewBufferString(query)),
	)

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to read from upstream")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unable to read from upstream")
	}

	projects, err := unmarshal(response.Body)

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to read response from upstream")
	}

	return projects, nil
}
