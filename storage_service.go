package storage

import (
	"context"
	. "github.com/common-go/storage"
)

type StorageService interface {
	Upload(ctx context.Context, contentImage File) (*StorageResult, error)
	Delete(ctx context.Context, fileName string) (bool, error)
}
