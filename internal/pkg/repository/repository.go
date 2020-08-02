package repository

import metav1 "mall/pkg/meta/v1"

// Repository is an interface for a repository.
type Repository interface {
	// Get returns the record for the given id.
	Get(id int) (record metav1.Resource, err error)
	// GetOne return one record filter by the conditions.
	GetOne(conditions interface{}) (record metav1.Resource, err error)
	// GetAll return all records filter by the conditions.
	GetAll(conditions interface{}) (records []metav1.Resource, err error)

	// Exist return one record does exist in table.
	Exist(conditions interface{}) (bool, error)

	// Create record.
	Create(record metav1.Resource) error
}
