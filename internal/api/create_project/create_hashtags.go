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

func createHashtags(
	db store.Database,
	args Args,
	project *models.Project,
	createdAt *time.Time,
) ([]*models.Hashtag, error) {
	hashTags := make([]*models.Hashtag, 0)

	for _, htArg := range args.Hashtags {
		errMsg := fmt.Sprintf("unable to insert hash tag '%s'", htArg)

		hashtag, err := store.Find(db, &models.Hashtag{Name: utils.PointerTo(htArg)})

		if errors.Is(err, store.ErrNotFound) {
			ht := &models.Hashtag{
				Name:      utils.PointerTo(htArg),
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

		hashTags = append(hashTags, hashtag)
	}

	return hashTags, nil
}
