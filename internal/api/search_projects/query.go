package search_projects

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/aakash-rajur/fold-se/internal/utils"
	"github.com/joomcode/errorx"
	"reflect"
	"strings"
	"text/template"
)

func inferQuery(args Args) (string, error) {
	if args.User != nil && len(*args.User) > 0 {
		return fieldEqQuery(
			fieldEqQueryArgs{
				FieldName: "users.name",
				Value:     *args.User,
			},
		)
	}

	if args.Hashtag != nil && len(*args.Hashtag) > 0 {
		return fieldEqQuery(
			fieldEqQueryArgs{
				FieldName: "hashtags.name",
				Value:     *args.Hashtag,
			},
		)
	}

	if args.Fuzziness != nil && *args.Fuzziness > 0 {
		return fuzzyQuery(
			fuzzyQueryArgs{
				Fuzziness: *args.Fuzziness,
				Fields: []FuzzyField{
					{
						Name:  "slug",
						Value: *(utils.Or(args.Slug, utils.PointerTo(""))),
					},
					{
						Name:  "description",
						Value: *(utils.Or(args.Description, utils.PointerTo(""))),
					},
				},
			},
		)
	}

	return "", errors.New("unsupported query")
}

func fieldEqQuery(args fieldEqQueryArgs) (string, error) {
	return buildQuery("eq", fieldEqQueryTmpl, args)
}

type fieldEqQueryArgs struct {
	FieldName string `json:"fieldName,omitempty"`
	Value     string `json:"value,omitempty"`
}

//go:embed field-eq.json.tmpl
var fieldEqQueryTmpl string

func fuzzyQuery(args fuzzyQueryArgs) (string, error) {
	return buildQuery("fuzzy", fuzzyQueryTmpl, args)
}

type fuzzyQueryArgs struct {
	Fuzziness int          `json:"fuzziness,omitempty"`
	Fields    []FuzzyField `json:"fields,omitempty"`
}

type FuzzyField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//go:embed fuzzy.json.tmpl
var fuzzyQueryTmpl string

func buildQuery[A any](name string, content string, args A) (string, error) {
	helpers := template.FuncMap{
		"isLast": func(index int, array interface{}) bool {
			return index == reflect.ValueOf(array).Len()-1
		},
		"ToUpper": strings.ToUpper,
		"ToJson": func(v interface{}) string {
			b, _ := json.Marshal(v)

			return string(b)
		},
	}

	tmpl, err := template.New(name).Funcs(helpers).Parse(content)

	if err != nil {
		return "", errorx.IllegalFormat.Wrap(err, "unable to parse model template")
	}

	var fuzzyBuffer bytes.Buffer

	err = tmpl.Execute(&fuzzyBuffer, args)

	if err != nil {
		return "", errorx.InternalError.Wrap(err, "unable to build query template")
	}

	query := fuzzyBuffer.String()

	return query, err
}
