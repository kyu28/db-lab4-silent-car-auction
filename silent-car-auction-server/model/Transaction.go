package model

import (
    "fmt"
    "database/sql"
)

// Data
type Transaction struct {
    TransactionID uint32
    AutoID uint32
    CustomerID uint32
    Bid float64
    BidDate string // date
}

// Implements Table
func (transaction Transaction) Columns() []string {
    return []string {"auto_id", "customer_id", "bid", "bid_date"}
}

func (transaction Transaction) Name() string {
    return "tb_transaction"
}

func (transaction Transaction) ToJSONArray(rows *sql.Rows) ([]map[string]interface{}, error) {
    array := make([]map[string]interface{}, 0)
    for rows.Next() {
        var t Transaction
        err := rows.Scan(&t.TransactionID, &t.AutoID, &t.CustomerID, &t.Bid, &t.BidDate)
        if err != nil {
            return nil, err
        }
        t.BidDate = t.BidDate[:10]
        array = append(array, map[string]interface{} {
            "transaction_id": t.TransactionID,
            "auto_id": t.AutoID,
            "customer_id": t.CustomerID,
            "bid": fmt.Sprintf("$%.2f", t.Bid),
            "bid_date": t.BidDate,
        })
    }
    return array, nil
}