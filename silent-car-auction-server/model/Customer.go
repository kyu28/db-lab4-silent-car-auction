package model

import (
    "database/sql"
)

// Data
type Customer struct {
    FirstName string
    LastName string
    Location string
    Email string
}

// Implements Table
func (customer Customer) Columns() []string {
    return []string {"first_name", "last_name", "location", "email"}
}

func (customer Customer) Name() string {
    return "tb_customer"
}

func (customer Customer) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var c Customer
        err := rows.Scan(&c.FirstName, &c.LastName, &c.Location, &c.Email)
        if err != nil {
            return nil, err
        }
        array = append(array, map[string]interface{} {
            "first_name": c.FirstName,
            "last_name": c.LastName,
            "location": c.Location,
            "email": c.Email,
        })
    }
    return array, nil
}
