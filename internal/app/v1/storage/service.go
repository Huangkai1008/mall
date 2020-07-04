package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"mall/internal/pkg/constant"
)

type Service struct {
	logger   *zap.Logger
	minioCli *minio.Client
}

func NewService(logger *zap.Logger, minioCli *minio.Client) *Service {
	return &Service{
		logger:   logger.With(zap.String("type", "StorageService")),
		minioCli: minioCli,
	}
}

func (s *Service) PutObject(objectName string, fh *multipart.FileHeader) (*ObjectSchema, error) {
	ctx := context.Background()

	// Check to see if we already own this bucket, if not exists, make new bucket.
	exists, err := s.minioCli.BucketExists(ctx, constant.BucketName)
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioCheckBucketExistError)
	}
	if !exists {
		err = s.minioCli.MakeBucket(ctx, constant.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, errors.Wrap(err, constant.MinioMakeBucketError)
		}
		s.logger.Info(constant.MinioMakeBucketOk, zap.String("bucketName", constant.BucketName))
	}

	// Upload the file
	file, err := fh.Open()
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioReadFileError)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			s.logger.Warn(constant.CloseFileError, zap.Error(err))
		}
	}(file)

	// Get file mimetype
	mType, err := mimetype.DetectReader(file)
	if err != nil {
		return nil, errors.Wrap(err, constant.GetFileMimetypeError)
	}

	info, err := s.minioCli.PutObject(
		ctx,
		constant.BucketName,
		objectName,
		file,
		-1,
		minio.PutObjectOptions{ContentType: fmt.Sprintln(mType)},
	)
	if err != nil {
		return nil, errors.Wrap(err, constant.MinioPutObjectError)
	}
	return &ObjectSchema{
		Url:          s.getUrl(constant.BucketName, objectName),
		ETag:         info.ETag,
		LastModified: info.LastModified,
		Location:     info.Location,
		VersionID:    info.VersionID,
	}, nil
}

func (s *Service) delimiter() string {
	return "/"
}

// getDirName Get directory name as object name prefix.
func (s *Service) getDirName() string {
	return time.Now().Format("2006-01-02")
}

// getObjectName returns generate object name.
func (s *Service) getObjectName(objectName string) string {
	return s.getDirName() + s.delimiter() + objectName
}

// getUrl returns the uploaded file url.
func (s Service) getUrl(bucketName, objectName string) string {
	return fmt.Sprintln(s.minioCli.EndpointURL()) + bucketName + objectName
}
