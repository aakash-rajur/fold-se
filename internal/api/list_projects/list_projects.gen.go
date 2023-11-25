package list_projects

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/store"
	"strings"
	"time"
)

type ListProjectsArgs struct {
	Limit  *int64 `db:"limit" json:"limit"`
	Offset *int64 `db:"offset" json:"offset"`
}

func (args *ListProjectsArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Limit: %v", *args.Limit),
			fmt.Sprintf("Offset: %v", *args.Offset),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectsArgs{%s}", content)
}

func (args *ListProjectsArgs) Query(db store.Database) ([]*ListProjectsResult, error) {
	return store.Query[ListProjectsResult](db, args)
}

func (args *ListProjectsArgs) Sql() string {
	return listProjectsSql
}

type ListProjectsResult struct {
	CreatedAt   *time.Time       `db:"created_at" json:"created_at"`
	Description *string          `db:"description" json:"description"`
	Hashtags    *store.JsonArray `db:"hashtags" json:"hashtags"`
	Id          *int64           `db:"id" json:"id"`
	Name        *string          `db:"name" json:"name"`
	RecordCount *int64           `db:"record_count" json:"record_count"`
	Slug        *string          `db:"slug" json:"slug"`
	Users       *store.JsonArray `db:"users" json:"users"`
}

func (result *ListProjectsResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("CreatedAt: %v", *result.CreatedAt),
			fmt.Sprintf("Description: %v", *result.Description),
			fmt.Sprintf("Hashtags: %v", result.Hashtags),
			fmt.Sprintf("Id: %v", *result.Id),
			fmt.Sprintf("Name: %v", *result.Name),
			fmt.Sprintf("RecordCount: %v", *result.RecordCount),
			fmt.Sprintf("Slug: %v", *result.Slug),
			fmt.Sprintf("Users: %v", result.Users),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectsResult{%s}", content)
}

//go:embed list-projects.sql
var listProjectsSql string
