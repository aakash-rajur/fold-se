package store

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joomcode/errorx"
)

func Connect(env map[string]string) (*sqlx.DB, func(), error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env["POSTGRESQL_USERNAME"],
		env["POSTGRESQL_PASSWORD"],
		env["POSTGRESQL_HOST"],
		env["POSTGRESQL_PORT"],
		env["POSTGRESQL_DATABASE"],
		env["POSTGRESQL_SSLMODE"],
	)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, nil, err
	}

	disconnect := func() {
		_ = db.Close()
	}

	return db, disconnect, err
}

func GetDb(ctx context.Context) (*sqlx.DB, error) {
	value := ctx.Value(dbKey)

	db, ok := value.(*sqlx.DB)

	if !ok {
		return nil, errorx.AssertionFailed.New("unable to get db from context")
	}

	return db, nil
}

func WithDb(db *sqlx.DB, ctx context.Context) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

const dbKey = "DB"
