package repository

import (
    "fmt"
    "time"
    "silent-car-auction/model/function"
	"github.com/gin-gonic/gin"
)

// Function callFunction calls a function with the given name and arguments and returns the result as a JSON object array
// Only use function when the function returns results
func callFunction(table function.Result, function string, args []interface{}, c *gin.Context) {
    err := ensureDatabase()
    if err != nil {
        c.String(500, err.Error())
        return
    }
    sql := fmt.Sprintf("SELECT * FROM %v(", function)
    if args != nil && len(args) > 0 {
        for _, arg := range args {
            switch arg.(type) {
            case string:
                sql += fmt.Sprintf("'%v', ", arg)
            case int:
                sql += fmt.Sprintf("%v, ", arg)
            case time.Time:
                sql += fmt.Sprintf("TO_DATE('%v', 'YYYY-MM-DD'), ", arg)
            }
        }
        sql = sql[:len(sql) - 2]
    }
    sql += ");"
    fmt.Println(sql)
    rows, err := db.Query(sql)
    if err != nil {
        c.String(400, err.Error())
        return
    }
    jsonArray, err := table.ToJSONArray(rows)
    if err != nil {
        c.String(400, err.Error())
        return
    }
    c.JSON(200, jsonArray)
}