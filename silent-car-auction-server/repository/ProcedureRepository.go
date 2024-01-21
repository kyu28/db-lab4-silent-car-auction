package repository

import (
    "fmt"
    "time"
	"github.com/gin-gonic/gin"
)

// Function callProcedure calls a stored procedure with the given name and arguments
func callProcedure(procedure string, args []interface{}, c *gin.Context) {
    err := ensureDatabase()
    if err != nil {
        c.String(500, err.Error())
        return
    }
    sql := fmt.Sprintf("CALL %v(", procedure)
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
    execSQL(sql, c)
}