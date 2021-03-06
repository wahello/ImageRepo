// Code generated by entc, DO NOT EDIT.

package imagerepo

const (
	// Label holds the string label denoting the imagerepo type in the database.
	Label = "image_repo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"

	// Table holds the table name of the imagerepo in the database.
	Table = "image_repos"
	// ImagesTable is the table the holds the images relation/edge. The primary key declared below.
	ImagesTable = "image_repo_images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
)

// Columns holds all SQL columns for imagerepo fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ImagesPrimaryKey and ImagesColumn2 are the table columns denoting the
	// primary key for the images relation (M2M).
	ImagesPrimaryKey = []string{"image_repo_id", "image_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
