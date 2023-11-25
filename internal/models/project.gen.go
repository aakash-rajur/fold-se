package models

import (
	"fmt"
	"strings"
	"time"
)

type Project struct {
	Id          *int64     `db:"id" json:"id"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	Description *string    `db:"description" json:"description"`
	Name        *string    `db:"name" json:"name"`
	Slug        *string    `db:"slug" json:"slug"`
}

func (project *Project) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *project.Id),
			fmt.Sprintf("CreatedAt: %v", *project.CreatedAt),
			fmt.Sprintf("Description: %v", *project.Description),
			fmt.Sprintf("Name: %v", *project.Name),
			fmt.Sprintf("Slug: %v", *project.Slug),
		},
		", ",
	)

	return fmt.Sprintf("Project{%s}", content)
}

func (_ *Project) TableName() string {
	return "public.projects"
}

func (_ *Project) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ *Project) InsertQuery() string {
	return projectInsertSql
}

func (_ *Project) UpdateQuery() string {
	return projectUpdateSql
}

func (_ *Project) FindQuery() string {
	return projectFindSql
}

func (_ *Project) FindAllQuery() string {
	return projectFindAllSql
}

func (_ *Project) DeleteQuery() string {
	return projectDeleteSql
}

// language=postgresql
var projectInsertSql = `
INSERT INTO public.projects(
  created_at,
  description,
  name,
  slug
)
VALUES (
  :created_at,
  :description,
  :name,
  :slug
)
RETURNING
  id,
  created_at,
  description,
  name,
  slug;
`

// language=postgresql
var projectUpdateSql = `
UPDATE public.projects
SET
  id = :id,
  created_at = :created_at,
  description = :description,
  name = :name,
  slug = :slug
WHERE TRUE
  AND id = :id
RETURNING
  id,
  created_at,
  description,
  name,
  slug;
`

// language=postgresql
var projectFindSql = `
SELECT
  id,
  created_at,
  description,
  name,
  slug
FROM public.projects
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or id = :id)
  AND (CAST(:description AS TEXT) IS NULL or id = :id)
  AND (CAST(:name AS TEXT) IS NULL or id = :id)
  AND (CAST(:slug AS TEXT) IS NULL or id = :id)
LIMIT 1;
`

// language=postgresql
var projectFindAllSql = `
SELECT
  id,
  created_at,
  description,
  name,
  slug
FROM public.projects
WHERE TRUE
	AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or id = :id)
  AND (CAST(:description AS TEXT) IS NULL or id = :id)
  AND (CAST(:name AS TEXT) IS NULL or id = :id)
  AND (CAST(:slug AS TEXT) IS NULL or id = :id)
`

// language=postgresql
var projectDeleteSql = `
DELETE FROM public.projects
WHERE TRUE
  AND id = :id;
`
