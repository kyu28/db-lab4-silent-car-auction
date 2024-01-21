package function

import (
    "database/sql"
)

// Data
type Brand struct {
    AutoID uint32
    Location string
    TypeOfAuto string
    Mileage int32
    Year string // date
}

// Implements Result
func (brand Brand) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var b Brand
        err := rows.Scan(&b.AutoID, &b.Location, &b.TypeOfAuto, &b.Mileage, &b.Year)
        if err != nil {
            return nil, err
        }
        b.Year = b.Year[:4]
        array = append(array, map[string]interface{} {
            "auto_id": b.AutoID,
            "location": b.Location,
            "type_of_auto": b.TypeOfAuto,
            "mileage": b.Mileage,
            "year": b.Year,
        })
    }
    return array, nil
}