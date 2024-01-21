// Package model contains the structs that represent the tables in the database
package model

import (
	"database/sql"
)

type Table interface {

	// Function Columns returns a slice of strings representing only the MUTABLE, INSERTABLE columns' names of the table
	Columns() []string

	// Function Name returns the name of the table
	Name() string

	// Function ToJSONArray converts a sql.Rows object to a slice of maps to a JSON object array
	ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error)
}