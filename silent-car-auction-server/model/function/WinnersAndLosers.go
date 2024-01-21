package function

import (
    "fmt"
    "database/sql"
)

// Data
type WinnersAndLosers struct {
    AutoID uint32
    LastName string
    Bid float32
    MaximumBid float32
    IsWinner bool
}

// Implements Result
func (wal WinnersAndLosers) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var w WinnersAndLosers
        err := rows.Scan(&w.AutoID, &w.LastName, &w.Bid, &w.MaximumBid, &w.IsWinner)
        if err != nil {
            return nil, err
        }
        array = append(array, map[string]interface{} {
            "auto_id": w.AutoID,
            "last_name": w.LastName,
            "bid": fmt.Sprintf("$%.2f", w.Bid),
            "maximum_bid": fmt.Sprintf("$%.2f", w.MaximumBid),
            "is_winner": w.IsWinner,
        })
    }
    return array, nil
}