package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"

	"mall/internal/pkg/config"
	"mall/internal/pkg/constant"
)

type Options struct {
	*config.Config
}

// New returns new minioClient instance with options.
func New(opts *Options) (*minio.Client, error) {
	minioClient, err := minio.New(opts.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(opts.AccessKeyID, opts.SecretAccessKey, ""),
		Secure: opts.UseSSL,
		Region: opts.Region,
	})
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioConfigError)
	}
	return minioClient, err
}
