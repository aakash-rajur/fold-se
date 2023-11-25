package models

import (
	"fmt"
	"strings"
)

type ProjectHashtag struct {
	HashtagId *int64 `db:"hashtag_id" json:"hashtag_id"`
	ProjectId *int64 `db:"project_id" json:"project_id"`
}

func (projectHashtag *ProjectHashtag) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("HashtagId: %v", *projectHashtag.HashtagId),
			fmt.Sprintf("ProjectId: %v", *projectHashtag.ProjectId),
		},
		", ",
	)

	return fmt.Sprintf("ProjectHashtag{%s}", content)
}

func (_ *ProjectHashtag) TableName() string {
	return "public.project_hashtags"
}

func (_ *ProjectHashtag) PrimaryKey() []string {
	return []string{
		"hashtag_id",
		"project_id",
	}
}

func (_ *ProjectHashtag) InsertQuery() string {
	return projectHashtagInsertSql
}

func (_ *ProjectHashtag) UpdateQuery() string {
	return projectHashtagUpdateSql
}

func (_ *ProjectHashtag) FindQuery() string {
	return projectHashtagFindSql
}

func (_ *ProjectHashtag) FindAllQuery() string {
	return projectHashtagFindAllSql
}

func (_ *ProjectHashtag) DeleteQuery() string {
	return projectHashtagDeleteSql
}

// language=postgresql
var projectHashtagInsertSql = `
INSERT INTO public.project_hashtags(
  hashtag_id,
  project_id
)
VALUES (
  :hashtag_id,
  :project_id
)
RETURNING
  hashtag_id,
  project_id;
`

// language=postgresql
var projectHashtagUpdateSql = `
UPDATE public.project_hashtags
SET
  hashtag_id = :hashtag_id,
  project_id = :project_id
WHERE TRUE
  AND hashtag_id = :hashtag_id
  AND project_id = :project_id
RETURNING
  hashtag_id,
  project_id;
`

// language=postgresql
var projectHashtagFindSql = `
SELECT
  hashtag_id,
  project_id
FROM public.project_hashtags
WHERE TRUE
  AND (CAST(:hashtag_id AS INT8) IS NULL or hashtag_id = :hashtag_id)
  AND (CAST(:project_id AS INT8) IS NULL or project_id = :project_id)
LIMIT 1;
`

// language=postgresql
var projectHashtagFindAllSql = `
SELECT
  hashtag_id,
  project_id
FROM public.project_hashtags
WHERE TRUE
  AND (CAST(:hashtag_id AS INT8) IS NULL or hashtag_id = :hashtag_id)
  AND (CAST(:project_id AS INT8) IS NULL or project_id = :project_id);
`

// language=postgresql
var projectHashtagDeleteSql = `
DELETE FROM public.project_hashtags
WHERE TRUE
  AND hashtag_id = :hashtag_id
  AND project_id = :project_id;
`
