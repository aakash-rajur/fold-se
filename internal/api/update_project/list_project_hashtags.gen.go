package update_project

import (
	_ "embed"
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/store"
	"strings"
	"time"
)

type ListProjectHashtagsArgs struct {
	ProjectId *int64 `db:"project_id" json:"project_id"`
}

func (args *ListProjectHashtagsArgs) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("ProjectId: %v", *args.ProjectId),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectHashtagsArgs{%s}", content)
}

func (args *ListProjectHashtagsArgs) Query(db store.Database) ([]*ListProjectHashtagsResult, error) {
	return store.Query[ListProjectHashtagsResult](db, args)
}

func (args *ListProjectHashtagsArgs) Sql() string {
	return listProjectHashtagsSql
}

type ListProjectHashtagsResult struct {
	HashtagCreatedAt *time.Time `db:"hashtag_created_at" json:"hashtag_created_at"`
	HashtagId        *int64     `db:"hashtag_id" json:"hashtag_id"`
	HashtagName      *string    `db:"hashtag_name" json:"hashtag_name"`
	ProjectId        *int64     `db:"project_id" json:"project_id"`
}

func (result *ListProjectHashtagsResult) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("HashtagCreatedAt: %v", *result.HashtagCreatedAt),
			fmt.Sprintf("HashtagId: %v", *result.HashtagId),
			fmt.Sprintf("HashtagName: %v", *result.HashtagName),
			fmt.Sprintf("ProjectId: %v", *result.ProjectId),
		},
		", ",
	)

	return fmt.Sprintf("ListProjectHashtagsResult{%s}", content)
}

//go:embed list-project-hashtags.sql
var listProjectHashtagsSql string
