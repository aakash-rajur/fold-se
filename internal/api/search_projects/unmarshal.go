package search_projects

import (
	"encoding/json"
	"errors"
	lp "github.com/aakash-rajur/fold-se/internal/api/list_projects"
	"github.com/joomcode/errorx"
	"io"
)

func unmarshal(reader io.Reader) ([]lp.ListProjectsResult, error) {
	buffer, err := io.ReadAll(reader)

	if err != nil {
		return nil, errorx.InternalError.Wrap(err, "unable to read response from upstream")
	}

	result := make(map[string]interface{})

	err = json.Unmarshal(buffer, &result)

	if err != nil {
		return nil, err
	}

	hits, ok := result["hits"].(map[string]any)

	if !ok {
		return nil, errors.New("unable to parse result for elastic search")
	}

	results, ok := hits["hits"].([]any)

	if !ok {
		return nil, errors.New("unable to parse result for elastic search")
	}

	projects := make([]lp.ListProjectsResult, 0)

	for _, eac := range results {
		hit, ok := eac.(map[string]any)

		if !ok {
			continue
		}

		source, ok := hit["_source"]

		if !ok {
			continue
		}

		buffer, err := json.Marshal(source)

		if err != nil {
			continue
		}

		project := lp.ListProjectsResult{}

		err = json.Unmarshal(buffer, &project)

		if err != nil {
			continue
		}

		projects = append(projects, project)
	}

	return projects, nil
}
