package storage

import (
	"context"
	"github.com/WildEgor/e-shop-cdn/internal/configs"
	domains "github.com/WildEgor/e-shop-cdn/internal/domain"
	"io"
)

var _ IFileStorage = (*StorageAdapter)(nil)

// StorageAdapter for s3
type StorageAdapter struct {
	provider IFileStorage `wire:"-"`
}

// NewStorageAdapter create storage (acts like factory)
func NewStorageAdapter(config *configs.StorageConfig) *StorageAdapter {
	switch config.Type {
	case S3Provider:
		provider, _ := NewS3Storage(&S3StorageConfig{
			Endpoint:   config.Minio.Endpoint,
			Region:     config.Minio.Region,
			Bucket:     config.Minio.Bucket,
			AccessId:   config.Minio.AccessKey,
			Secret:     config.Minio.Secret,
			DownloadFn: config.DownloadUrl,
		})
		return &StorageAdapter{
			provider: provider,
		}
	default:
		return nil
	}
}

// DownloadURL
func (s *StorageAdapter) DownloadURL(objectName string) string {
	return s.provider.DownloadURL(objectName)
}

// Metadata
func (s *StorageAdapter) Metadata(objectName string) (*domains.FileMetadata, error) {
	return s.provider.Metadata(objectName)
}

// Upload
func (s *StorageAdapter) Upload(ctx context.Context, objectName string, reader io.Reader) error {
	return s.provider.Upload(ctx, objectName, reader)
}

// Download
func (s *StorageAdapter) Download(ctx context.Context, objectName string) ([]byte, error) {
	return s.provider.Download(ctx, objectName)
}

// Delete
func (s *StorageAdapter) Delete(objectName string) error {
	return s.provider.Delete(objectName)
}

// Exists
func (s *StorageAdapter) Exists(ctx context.Context, objectName string) (bool, error) {
	return s.provider.Exists(ctx, objectName)
}

// GetProvider
func (s *StorageAdapter) GetProvider() IFileStorage {
	return s.provider
}

// SetProvider
func (s *StorageAdapter) SetProvider(provider IFileStorage) {
	s.provider = provider
}

// Ping
func (s *StorageAdapter) Ping() error {
	if ping, ok := s.provider.(IPinger); ok {
		return ping.Ping()
	}
	return nil
}
