package schema

import (
	"mime/multipart"
	"time"

	"mall/internal/pkg/schema"
)

// ObjectCreateSchema is the object create schema.
type ObjectCreateSchema struct {
	schema.BaseSchema
	File *multipart.FileHeader `form:"file" binding:"required,file"`
}

// ObjectSchema is the object return schema.
type ObjectSchema struct {
	Url          string    `json:"url"`           // file url
	ETag         string    `json:"etag"`          // ETag message
	LastModified time.Time `json:"last_modified"` // last modified time
	Location     string    `json:"location"`      // location
	VersionID    string    `json:"version_id"`    // version id
}
