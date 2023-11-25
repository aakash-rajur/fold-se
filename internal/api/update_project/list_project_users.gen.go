package update_project

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/store"
	"strings"
	"time"
)

type ListProjectUsersArgs struct {
	ProjectId *int64 `db:"project_id" json:"project_id"`
}

func (args *ListProjectUsersArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("ProjectId: %v", *args.ProjectId),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectUsersArgs{%s}", content)
}

func (args *ListProjectUsersArgs) Query(db store.Database) ([]*ListProjectUsersResult, error) {
	return store.Query[ListProjectUsersResult](db, args)
}

func (args *ListProjectUsersArgs) Sql() string {
	return listProjectUsersSql
}

type ListProjectUsersResult struct {
	ProjectId     *int64     `db:"project_id" json:"project_id"`
	UserCreatedAt *time.Time `db:"user_created_at" json:"user_created_at"`
	UserId        *int64     `db:"user_id" json:"user_id"`
	UserName      *string    `db:"user_name" json:"user_name"`
}

func (result *ListProjectUsersResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("ProjectId: %v", *result.ProjectId),
			fmt.Sprintf("UserCreatedAt: %v", *result.UserCreatedAt),
			fmt.Sprintf("UserId: %v", *result.UserId),
			fmt.Sprintf("UserName: %v", *result.UserName),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectUsersResult{%s}", content)
}

//go:embed list-project-users.sql
var listProjectUsersSql string
