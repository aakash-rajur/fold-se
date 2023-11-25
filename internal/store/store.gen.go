package store

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"
)

func Insert[T model[P], P any](db Database, instances ...T) error {
	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.InsertQuery(), instance)

		if err != nil {
			return err
		}

		hasNext := rows.Next()

		if !hasNext {
			return ErrNotFound
		}

		err = rows.StructScan(instance)

		if err != nil {
			return err
		}

		err = rows.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func Update[T model[P], P any](db Database, instances ...T) error {
	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.UpdateQuery(), instance)

		if err != nil {
			return err
		}

		rows.Next()

		err = rows.StructScan(instance)

		_ = rows.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func Find[T model[P], P any](db Database, instance T) (T, error) {
	result := new(P)

	rows, err := db.NamedQuery(instance.FindQuery(), *instance)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	hasOne := rows.Next()

	if !hasOne {
		return nil, ErrNotFound
	}

	err = rows.StructScan(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindMany[T model[P], P any](db Database, instance T) ([]T, error) {
	rows, err := db.NamedQuery(instance.FindAllQuery(), *instance)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	result := make([]T, 0)

	for rows.Next() {
		var instance T

		err = rows.StructScan(instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	return result, nil
}

func Delete[T model[P], P any](db Database, instance T) error {
	_, err := db.NamedExec(instance.DeleteQuery(), instance)

	return err
}

type model[P any] interface {
	*P

	TableName() string

	PrimaryKey() []string

	InsertQuery() string

	UpdateQuery() string

	FindQuery() string

	FindAllQuery() string

	DeleteQuery() string
}

func Query[R any](db Database, args queryable) ([]*R, error) {
	re, err := regexp.Compile(`-{2,}\s*([\w\W\s\S]*?)(\n|\z)`)

	if err != nil {
		return nil, err
	}

	query := re.ReplaceAllString(args.Sql(), "$2")

	rows, err := db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	result := make([]*R, 0)

	for rows.Next() {
		instance := new(R)

		err = rows.StructScan(instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	return result, nil
}

type queryable interface {
	Sql() string
}

type Database interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type JsonObject map[string]interface{}

func (j *JsonObject) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type JsonArray []map[string]interface{}

func (j *JsonArray) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

var ErrNotFound = errors.New("entity not found in database")
