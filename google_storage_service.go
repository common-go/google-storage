package storage

import (
	"cloud.google.com/go/storage"
	. "github.com/common-go/storage"
	"golang.org/x/net/context"
	"path"
)

const storageUrl = "https://storage.googleapis.com"

type GoogleStorageService struct {
	Client *storage.Client
	Config Config
	Bucket *storage.BucketHandle
}

func NewGoogleStorageService(client *storage.Client, config Config) *GoogleStorageService {
	return &GoogleStorageService{client,
		config,
		client.Bucket(config.BucketName)}
}

func (s GoogleStorageService) Upload(ctx context.Context, objectFile File) (*StorageResult, error) {
	dir := objectFile.FileName
	if len(s.Config.SubDirectory) > 0 {
		dir = path.Join(s.Config.SubDirectory, objectFile.FileName)
	}
	object := s.Bucket.Object(dir)
	wc := object.NewWriter(ctx)
	wc.ContentType = "image/png"

	if len(objectFile.ContentType) > 0 {
		wc.ContentType = objectFile.ContentType
	}
	if _, err := wc.Write(objectFile.BytesData); err != nil {
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}
	if s.Config.PermissionFileRoleAll {
		if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, err
		}
	}
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return &StorageResult{Status: 1, Name: objectFile.FileName, MediaLink: attrs.MediaLink, Link: getLinkPublic(s.Config.BucketName, dir)}, nil
}

func getLinkPublic(bucketName string, remoteFile string) string {
	return path.Join(storageUrl, bucketName, remoteFile)
}

func (s GoogleStorageService) Delete(ctx context.Context, fileName string) (bool, error) {
	dir := fileName
	if len(s.Config.SubDirectory) > 0 {
		dir = path.Join(s.Config.SubDirectory, fileName)
	}
	if err := s.Bucket.Object(dir).Delete(ctx); err != nil {
		return false, err
	}
	return true, nil
}
