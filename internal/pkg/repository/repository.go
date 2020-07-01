package repository

type Repository interface {
	// GetAll return all records filter by the conditions.
	GetAll(conditions interface{}) (records []*interface{}, err error)
	// GetOne return one record filter by the conditions.
	GetOne(conditions interface{}) (record *interface{}, err error)

	// Exist return one record does exist in table.
	Exist(conditions interface{}) (bool, error)

	// Create record.
	Create(record interface{}) error
}
