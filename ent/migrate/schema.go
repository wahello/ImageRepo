// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "file_location", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "public", Type: field.TypeBool},
		{Name: "user_images", Type: field.TypeInt, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "images_users_images",
				Columns: []*schema.Column{ImagesColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImageReposColumns holds the columns for the "image_repos" table.
	ImageReposColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// ImageReposTable holds the schema information for the "image_repos" table.
	ImageReposTable = &schema.Table{
		Name:        "image_repos",
		Columns:     ImageReposColumns,
		PrimaryKey:  []*schema.Column{ImageReposColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ImageRepoImagesColumns holds the columns for the "image_repo_images" table.
	ImageRepoImagesColumns = []*schema.Column{
		{Name: "image_repo_id", Type: field.TypeInt},
		{Name: "image_id", Type: field.TypeInt},
	}
	// ImageRepoImagesTable holds the schema information for the "image_repo_images" table.
	ImageRepoImagesTable = &schema.Table{
		Name:       "image_repo_images",
		Columns:    ImageRepoImagesColumns,
		PrimaryKey: []*schema.Column{ImageRepoImagesColumns[0], ImageRepoImagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "image_repo_images_image_repo_id",
				Columns: []*schema.Column{ImageRepoImagesColumns[0]},

				RefColumns: []*schema.Column{ImageReposColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "image_repo_images_image_id",
				Columns: []*schema.Column{ImageRepoImagesColumns[1]},

				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ImagesTable,
		ImageReposTable,
		UsersTable,
		ImageRepoImagesTable,
	}
)

func init() {
	ImagesTable.ForeignKeys[0].RefTable = UsersTable
	ImageRepoImagesTable.ForeignKeys[0].RefTable = ImageReposTable
	ImageRepoImagesTable.ForeignKeys[1].RefTable = ImagesTable
}
