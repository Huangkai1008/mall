package storage

import (
	"mime/multipart"
	"time"

	"mall/internal/pkg/schema"
)

// ObjectCreateSchema is the object create schema
type ObjectCreateSchema struct {
	schema.BaseSchema
	File *multipart.FileHeader `form:"file" binding:"required,file"`
}

type ObjectSchema struct {
	Url          string    // file url
	ETag         string    `json:"etag"` // ETag message
	LastModified time.Time // last modified time
	Location     string    // location
	VersionID    string    `json:"version_id"` // version id
}
