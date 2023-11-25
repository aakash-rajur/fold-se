package es

import (
	"context"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/joomcode/errorx"
)

func Client(env map[string]string) (*es.Client, error) {
	apiEndpoint := env["ELASTICSEARCH_API_ENDPOINT"]

	cfg := es.Config{
		Addresses:           []string{apiEndpoint},
		CompressRequestBody: true,
	}

	return es.NewClient(cfg)
}

func GetEsc(ctx context.Context) (*es.Client, error) {
	value := ctx.Value(esKey)

	client, ok := value.(*es.Client)

	if !ok {
		return nil, errorx.AssertionFailed.New("unable to get es client from context")
	}

	return client, nil
}

func WithEs(client *es.Client, ctx context.Context) context.Context {
	return context.WithValue(ctx, esKey, client)
}

const esKey = "ES"
