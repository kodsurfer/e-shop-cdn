package storage

import (
	"context"
	domains "github.com/WildEgor/e-shop-cdn/internal/domain"
	"io"
)

const (
	/*S3Provider s3 type*/
	S3Provider string = "s3"
)

// IFileStorage naive for any storage
type IFileStorage interface {
	// Upload file to storage
	Upload(ctx context.Context, objectName string, reader io.Reader) error
	// Download return file bytes by path from storage
	Download(ctx context.Context, objectName string) ([]byte, error)
	// Delete file by path
	Delete(objectName string) error
	// Exists check if file exists by path
	Exists(ctx context.Context, objectName string) (bool, error)
	// Metadata return filemeta
	Metadata(objectName string) (*domains.FileMetadata, error)
	// DownloadURL create download url
	DownloadURL(objectName string) string
}

// IPinger for pings (for example, not suitable for fs)
type IPinger interface {
	Ping() error
}
