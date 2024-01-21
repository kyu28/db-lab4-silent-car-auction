package model

import (
    "database/sql"
)

// Data
type AutoIndex struct {
    AutoID uint32
    VIN string
}

// Implements Table
func (autoIndex AutoIndex) Columns() []string {
    return []string {"VIN"}
}

func (autoIndex AutoIndex) Name() string {
    return "tb_auto_index"
}

func (autoIndex AutoIndex) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var a AutoIndex
        err := rows.Scan(&a.AutoID, &a.VIN)
        if err != nil {
            return nil, err
        }
        array = append(array, map[string]interface{} {
            "auto_id": a.AutoID,
            "VIN": a.VIN,
        })
    }
    return array, nil
}