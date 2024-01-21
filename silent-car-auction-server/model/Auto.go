package model

import (
    "database/sql"
)

// Data
type Auto struct {
    VIN string
    Location string
    TypeOfAuto string
    Mileage int32
    Year string // date, year
}

// Implements Table
func (auto Auto) Columns() []string {
    return []string {"VIN", "location", "type_of_auto", "mileage", "year"}
}

func (auto Auto) Name() string {
    return "tb_auto"
}

func (auto Auto) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var a Auto
        err := rows.Scan(&a.VIN, &a.Location, &a.TypeOfAuto, &a.Mileage, &a.Year)
        if err != nil {
            return nil, err
        }
        a.Year = a.Year[:4]
        array = append(array, map[string]interface{} {
            "VIN": a.VIN,
            "location": a.Location,
            "type_of_auto": a.TypeOfAuto,
            "mileage": a.Mileage,
            "year": a.Year,
        })
    }
    return array, nil
}