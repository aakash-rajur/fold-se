package models

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id        *int64     `db:"id" json:"id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	Name      *string    `db:"name" json:"name"`
}

func (user *User) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Id: %v", *user.Id),
			fmt.Sprintf("CreatedAt: %v", *user.CreatedAt),
			fmt.Sprintf("Name: %v", *user.Name),
		},
		", ",
	)

	return fmt.Sprintf("User{%s}", content)
}

func (_ *User) TableName() string {
	return "public.users"
}

func (_ *User) PrimaryKey() []string {
	return []string{
		"id",
	}
}

func (_ *User) InsertQuery() string {
	return userInsertSql
}

func (_ *User) UpdateQuery() string {
	return userUpdateSql
}

func (_ *User) FindQuery() string {
	return userFindSql
}

func (_ *User) FindAllQuery() string {
	return userFindAllSql
}

func (_ *User) DeleteQuery() string {
	return userDeleteSql
}

// language=postgresql
var userInsertSql = `
INSERT INTO public.users(
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
var userUpdateSql = `
UPDATE public.users
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
var userFindSql = `
SELECT
  id,
  created_at,
  name
FROM public.users
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or created_at = :created_at)
  AND (CAST(:name AS TEXT) IS NULL or name = :name)
LIMIT 1;
`

// language=postgresql
var userFindAllSql = `
SELECT
  id,
  created_at,
  name
FROM public.users
WHERE TRUE
  AND (CAST(:id AS INT8) IS NULL or id = :id)
  AND (CAST(:created_at AS TIMESTAMPTZ) IS NULL or created_at = :created_at)
  AND (CAST(:name AS TEXT) IS NULL or name = :name);
`

// language=postgresql
var userDeleteSql = `
DELETE FROM public.users
WHERE TRUE
  AND id = :id;
`
