// Package function contains the entities that represent the tables return from functions in the database
package function

import (
	"database/sql"
)

type Result interface {
	// Function ToJSONArray converts a sql.Rows object to a slice of maps to a JSON object array
	ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error)
}