package model

import (
    "database/sql"
)

// Data
type Winner struct {
    AutoID uint32
    CustomerID uint32
}

// Implements Table
func (winner Winner) Columns() []string {
    return []string { "auto_id", "customer_id" }
}

func (winner Winner) Name() string {
    return "tb_winner"
}

func (winner Winner) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var w Winner
        err := rows.Scan(&w.AutoID, &w.CustomerID)
        if err != nil {
            return nil, err
        }
        array = append(array, map[string]interface{} {
            "auto_id": w.AutoID,
            "customer_id": w.CustomerID,
        })
    }
    return array, nil
}