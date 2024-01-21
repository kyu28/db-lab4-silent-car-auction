package function

import (
    "fmt"
    "database/sql"
)

// Data
type MaxBid struct {
    AutoID uint32
    Location string
    Year string // date
    MaximumBid float32
}

// Implements Result
func (mb MaxBid) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var m MaxBid
        err := rows.Scan(&m.AutoID, &m.Location, &m.Year, &m.MaximumBid)
        if err != nil {
            return nil, err
        }
        m.Year = m.Year[:4]
        array = append(array, map[string]interface{} {
            "auto_id": m.AutoID,
            "location": m.Location,
            "year": m.Year,
            "maximum_bid": fmt.Sprintf("$%.2f", m.MaximumBid), 
        })
    }
    return array, nil
}