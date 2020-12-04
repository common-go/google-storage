package storage

import (
	"cloud.google.com/go/storage"
	"context"
	. "github.com/common-go/storage"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, config Config) (*storage.Client, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.CredentialsFile)) //"resource/key.json"
	if err != nil {
		return nil, err
	}
	return client, nil
}
