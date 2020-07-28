package v1

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	Kind string `json:"kind,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	APIVersion string `json:"api_version,omitempty"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// ObjectMeta is also used by gorm.
type ObjectMeta struct {
	// ID is the unique in time and space value for this object. It is typically generated by
	// the storage on successful creation of a resource and is not allowed to change on PUT
	// operations.
	//
	// Populated by the system.
	// Read-only.
	ID uint64 `json:"id,omitempty" gorm:"type:bigint(11) UNSIGNED AUTO_INCREMENT;primaryKey"`

	// CreatedAt is a timestamp representing the server time when this object was created.
	// It is represented in `2006-01-02 15:04:05` format.
	//
	// Populated by the system.
	// Read-only.
	CreatedAt JsonTime `json:"create_time,omitempty" gorm:"type:datetime;column:create_time;comment:创建时间"`

	// UpdatedAt is a timestamp representing the server time when this object was updated.
	// It is represented in `2006-01-02 15:04:05` format.
	//
	// Populated by the system.
	// Read-only.
	UpdatedAt JsonTime `json:"update_time,omitempty" gorm:"type:datetime;column:update_time;comment:更新时间"`
}
