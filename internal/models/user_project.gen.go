package models

import (
	"fmt"
	"strings"
)

type UserProject struct {
	ProjectId *int64 `db:"project_id" json:"project_id"`
	UserId    *int64 `db:"user_id" json:"user_id"`
}

func (userProject *UserProject) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("ProjectId: %v", *userProject.ProjectId),
			fmt.Sprintf("UserId: %v", *userProject.UserId),
		},
		", ",
	)

	return fmt.Sprintf("UserProject{%s}", content)
}

func (_ *UserProject) TableName() string {
	return "public.user_projects"
}

func (_ *UserProject) PrimaryKey() []string {
	return []string{
		"project_id",
		"user_id",
	}
}

func (_ *UserProject) InsertQuery() string {
	return userProjectInsertSql
}

func (_ *UserProject) UpdateQuery() string {
	return userProjectUpdateSql
}

func (_ *UserProject) FindQuery() string {
	return userProjectFindSql
}

func (_ *UserProject) FindAllQuery() string {
	return userProjectFindAllSql
}

func (_ *UserProject) DeleteQuery() string {
	return userProjectDeleteSql
}

// language=postgresql
var userProjectInsertSql = `
INSERT INTO public.user_projects(
  project_id,
  user_id
)
VALUES (
  :project_id,
  :user_id
)
RETURNING
  project_id,
  user_id;
`

// language=postgresql
var userProjectUpdateSql = `
UPDATE public.user_projects
SET
  project_id = :project_id,
  user_id = :user_id
WHERE TRUE
  AND project_id = :project_id
  AND user_id = :user_id
RETURNING
  project_id,
  user_id;
`

// language=postgresql
var userProjectFindSql = `
SELECT
  project_id,
  user_id
FROM public.user_projects
WHERE TRUE
  AND (CAST(:project_id AS INT8) IS NULL or project_id = :project_id)
  AND (CAST(:user_id AS INT8) IS NULL or user_id = :user_id)
LIMIT 1;
`

// language=postgresql
var userProjectFindAllSql = `
SELECT
  project_id,
  user_id
FROM public.user_projects
WHERE TRUE
  AND (CAST(:project_id AS INT8) IS NULL or project_id = :project_id)
  AND (CAST(:user_id AS INT8) IS NULL or user_id = :user_id);
`

// language=postgresql
var userProjectDeleteSql = `
DELETE FROM public.user_projects
WHERE TRUE
  AND project_id = :project_id
  AND user_id = :user_id;
`
