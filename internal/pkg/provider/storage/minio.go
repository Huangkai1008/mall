package storage

import (
	"github.com/minio/minio-go/v7"

	minioCli "github.com/Huangkai1008/micro-kit/pkg/storage/minio"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// NewMinioClient returns new minio client instance.
func NewMinioClient(c *config.Config) (*minio.Client, error) {
	return minioCli.New(
		minioCli.WithEndpoint(c.Minio.Endpoint),
		minioCli.WithAccessKeyID(c.Minio.AccessKeyID),
		minioCli.WithUseSSL(c.Minio.UseSSL),
		minioCli.WithSecretAccessKey(c.Minio.SecretAccessKey),
		minioCli.WithRegion(c.Minio.Region),
	)
}
