// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TblOdoColumns holds the columns for the "tbl_odo" table.
	TblOdoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString},
		{Name: "done", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeInt},
	}
	// TblOdoTable holds the schema information for the "tbl_odo" table.
	TblOdoTable = &schema.Table{
		Name:       "tbl_odo",
		Columns:    TblOdoColumns,
		PrimaryKey: []*schema.Column{TblOdoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tbl_odo_tbl_user_todos",
				Columns:    []*schema.Column{TblOdoColumns[3]},
				RefColumns: []*schema.Column{TblUserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TblUserColumns holds the columns for the "tbl_user" table.
	TblUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "firstname", Type: field.TypeString, Nullable: true},
		{Name: "lastname", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString},
	}
	// TblUserTable holds the schema information for the "tbl_user" table.
	TblUserTable = &schema.Table{
		Name:       "tbl_user",
		Columns:    TblUserColumns,
		PrimaryKey: []*schema.Column{TblUserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TblOdoTable,
		TblUserTable,
	}
)

func init() {
	TblOdoTable.ForeignKeys[0].RefTable = TblUserTable
	TblOdoTable.Annotation = &entsql.Annotation{
		Table: "tbl_odo",
	}
	TblUserTable.Annotation = &entsql.Annotation{
		Table: "tbl_user",
	}
}
