package storage

import (
	"time"

	"mall/internal/pkg/schema"
)

// ObjectCreateSchema is the object create schema
type ObjectCreateSchema struct {
	schema.BaseSchema
	file string `form:"file" binding:"required,file"`
}

type ObjectSchema struct {
	Url          string    // file url
	ETag         string    `json:"etag"` // ETag message
	LastModified time.Time // last modified time
	Location     string    // location
	VersionID    string    `json:"version_id"` // version id
}
