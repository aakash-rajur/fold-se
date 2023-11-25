package models

import (
	"fmt"
	"strings"
	"time"
)

type Hashtag struct {
	Id        *int64     `db:"id" json:"id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	Name      *string    `db:"name" json:"name"`
}

func (hashtag *Hashtag) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *hashtag.Id),
			fmt.Sprintf("CreatedAt: %v", *hashtag.CreatedAt),
			fmt.Sprintf("Name: %v", *hashtag.Name),
		},
		", ",
	)

	return fmt.Sprintf("Hashtag{%s}", content)
}

func (_ *Hashtag) TableName() string {
	return "public.hashtags"
}

func (_ *Hashtag) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ *Hashtag) InsertQuery() string {
	return hashtagInsertSql
}

func (_ *Hashtag) UpdateQuery() string {
	return hashtagUpdateSql
}

func (_ *Hashtag) FindQuery() string {
	return hashtagFindSql
}

func (_ *Hashtag) FindAllQuery() string {
	return hashtagFindAllSql
}

func (_ *Hashtag) DeleteQuery() string {
	return hashtagDeleteSql
}

// language=postgresql
var hashtagInsertSql = `
INSERT INTO public.hashtags(
  created_at,
  name
)
VALUES (
  :created_at,
  :name
)
RETURNING
  id,
  created_at,
  name;
`

// language=postgresql
var hashtagUpdateSql = `
UPDATE public.hashtags
SET
  id = :id,
  created_at = :created_at,
  name = :name
WHERE TRUE
  AND id = :id
RETURNING
  id,
  created_at,
  name;
`

// language=postgresql
var hashtagFindSql = `
SELECT
  id,
  created_at,
  name
FROM public.hashtags
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or created_at = :created_at)
  AND (CAST(:name AS TEXT) IS NULL or name = :name)
LIMIT 1;
`

// language=postgresql
var hashtagFindAllSql = `
SELECT
  id,
  created_at,
  name
FROM public.hashtags
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or created_at = :created_at)
  AND (CAST(:name AS TEXT) IS NULL or name = :name);
`

// language=postgresql
var hashtagDeleteSql = `
DELETE FROM public.hashtags
WHERE TRUE
  AND id = :id;
`
