package minio

import (
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"

	"mall/internal/pkg/constant"
)

// New returns new minioClient instance with options.
func New(o *Options) (*minio.Client, error) {
	minioClient, err := minio.New(o.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(o.AccessKeyID, o.SecretAccessKey, ""),
		Secure: o.UseSSL,
		Region: o.Region,
	})
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioConfigError)
	}
	return minioClient, err
}

var ProviderSet = wire.NewSet(New, NewOptions)
