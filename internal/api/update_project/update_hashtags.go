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

func updateHashtags(
	db store.Database,
	args Args,
	project *models.Project,
	createdAt *time.Time,
) ([]*models.Hashtag, error) {
	lpha := &ListProjectHashtagsArgs{ProjectId: args.Id}

	projectHashtags, err := lpha.Query(db)

	if err != nil {
		return nil, err
	}

	projectHashtagMap := make(map[string]*ListProjectHashtagsResult)

	for _, ph := range projectHashtags {
		projectHashtagMap[*ph.HashtagName] = ph
	}

	projectHashtagArgsMap := make(map[string]bool)

	hashtags := make([]*models.Hashtag, 0)

	for _, projectHashtagArg := range args.Hashtags {
		// map to lookup for the next loop
		projectHashtagArgsMap[projectHashtagArg] = true

		_, ok := projectHashtagMap[projectHashtagArg]

		if ok {
			continue
		}

		errMsg := fmt.Sprintf("unable to insert hash tag '%s'", projectHashtagArg)

		hashtag, err := store.Find(db, &models.Hashtag{Name: utils.PointerTo(projectHashtagArg)})

		if errors.Is(err, store.ErrNotFound) {
			ht := &models.Hashtag{
				Name:      utils.PointerTo(projectHashtagArg),
				CreatedAt: createdAt,
			}

			err = store.Insert(db, ht)

			if err != nil {
				return nil, errorx.InternalError.Wrap(err, errMsg)
			}

			hashtag = ht
		} else if err != nil {
			return nil, errorx.InternalError.Wrap(err, errMsg)
		}

		hp := &models.ProjectHashtag{
			ProjectId: project.Id,
			HashtagId: hashtag.Id,
		}

		err = store.Insert(db, hp)

		if err != nil {
			return nil, errorx.InternalError.Wrap(err, "unable to associate hashtag with project")
		}

		hashtags = append(hashtags, hashtag)
	}

	for _, projectHashTag := range projectHashtags {
		_, ok := projectHashtagArgsMap[*projectHashTag.HashtagName]

		if ok {
			hashtag := &models.Hashtag{
				Id:        projectHashTag.HashtagId,
				CreatedAt: projectHashTag.HashtagCreatedAt,
				Name:      projectHashTag.HashtagName,
			}

			hashtags = append(hashtags, hashtag)

			continue
		}

		err := store.Delete(
			db,
			&models.ProjectHashtag{
				ProjectId: projectHashTag.ProjectId,
				HashtagId: projectHashTag.HashtagId,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return hashtags, nil
}
