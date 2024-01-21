package model

import (
    "database/sql"
)

// Data
type CustomerIndex struct {
    CustomerId uint32
    Email string
}

// Implements Table
func (customerIndex CustomerIndex) Columns() []string {
    return []string {"email"}
}

func (customerIndex CustomerIndex) Name() string {
    return "tb_customer_index"
}

func (customerIndex CustomerIndex) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var c CustomerIndex
        err := rows.Scan(&c.CustomerId, &c.Email)
        if err != nil {
            return nil, err
        }
        array = append(array, map[string]interface{} {
            "customer_id": c.CustomerId,
            "email": c.Email,
        })
    }
    return array, nil
}
