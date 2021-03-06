package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/Huangkai1008/mall/internal/app/v1/storage/schema"
	"github.com/Huangkai1008/mall/internal/pkg/config"
	"github.com/Huangkai1008/mall/internal/pkg/constant"
)

type StorageService struct {
	Region   string
	logger   *zap.Logger
	minioCli *minio.Client
}

func NewService(c *config.Config, logger *zap.Logger, minioCli *minio.Client) *StorageService {
	return &StorageService{
		logger:   logger.With(zap.String("type", "StorageService")),
		minioCli: minioCli,
		Region:   c.Minio.Region,
	}
}

// PutObject put object to bucket.
func (s *StorageService) PutObject(objectName string, fh *multipart.FileHeader) (*schema.ObjectSchema, error) {
	ctx := context.Background()

	// Check to see if we already own this bucket, if not exists, make new bucket.
	exists, err := s.minioCli.BucketExists(ctx, constant.BucketName)
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioCheckBucketExistError)
	}
	if !exists {
		err = s.minioCli.MakeBucket(ctx, constant.BucketName, minio.MakeBucketOptions{Region: s.Region})
		if err != nil {
			return nil, errors.Wrap(err, constant.MinioMakeBucketError)
		}

		err = s.SetReadOnlyBucketPolicy(ctx, constant.BucketName)
		if err != nil {
			return nil, errors.Wrap(err, constant.MinioSetPolicyError)
		}
		s.logger.Info(constant.MinioMakeBucketOk, zap.String("bucketName", constant.BucketName))
	}

	// Upload the file
	file, err := fh.Open()
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioReadFileError)
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			s.logger.Warn(constant.CloseFileError, zap.Error(err))
		}
	}(file)

	// Get file content-type
	contentType := fh.Header.Get("content-type")

	info, err := s.minioCli.PutObject(
		ctx,
		constant.BucketName,
		s.getObjectName(objectName),
		file,
		-1,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return nil, errors.WithMessage(err, constant.MinioPutObjectError)
	}
	return &schema.ObjectSchema{
		Url:          s.getUrl(constant.BucketName, info.Key),
		ETag:         info.ETag,
		LastModified: info.LastModified,
		Location:     info.Location,
		VersionID:    info.VersionID,
	}, nil
}

// SetReadOnlyBucketPolicy set read-only permissions on an existing bucket.
func (s *StorageService) SetReadOnlyBucketPolicy(ctx context.Context, bucketName string) error {
	policy := fmt.Sprintf(`
    {
		"Version": "2012-10-17",
        "Statement": [
            {
            	"Effect": "Allow",
				"Principal": {"AWS": "*"},
 				"Action": ["s3:GetBucketLocation", "s3:ListBucket"],
				"Resource": "arn:aws:s3:::%[1]s"
            },
			{
            	"Effect": "Allow",
				"Principal": {"AWS": "*"},
 				"Action": ["s3:GetObject"],
				"Resource": "arn:aws:s3:::%[1]s/*"
            }
        ]
	}`,
		bucketName,
	)
	if err := s.minioCli.SetBucketPolicy(ctx, bucketName, policy); err != nil {
		return errors.Wrap(err, constant.MinioSetPolicyError)
	} else {
		return nil
	}
}

func (s *StorageService) delimiter() string {
	return "/"
}

// getDirName Get directory name as object name prefix.
func (s *StorageService) getDirName() string {
	return time.Now().Format("2006-01-02")
}

// getObjectName returns generate object name.
func (s *StorageService) getObjectName(objectName string) string {
	return s.getDirName() + s.delimiter() + objectName
}

// getUrl returns the uploaded file url.
func (s *StorageService) getUrl(bucketName, objectName string) string {
	return fmt.Sprint(s.minioCli.EndpointURL()) + "/" + bucketName + "/" + objectName
}

var ProviderSet = wire.NewSet(NewService)
