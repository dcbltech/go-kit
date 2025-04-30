package storage

import "context"

type Storage interface {
	Exists(ctx context.Context, path, name string) (exists bool, err error)
	Store(ctx context.Context, path, name string, data []byte) error
	Load(ctx context.Context, path, name string) (data []byte, err error)
	Delete(ctx context.Context, path, name string) error
	DeleteAll(ctx context.Context, path string) error
	Copy(ctx context.Context, srcPath, srcName, dstPath, dstName string) error
	GenerateSignedURL(path, name string) (url string, err error)
}
