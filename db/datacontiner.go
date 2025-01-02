package db

// DataContainer is an abstraction for both SQL tables and NoSQL collections.
type DataContainer interface {
	Name() string // Returns the name of the table/collection
}
