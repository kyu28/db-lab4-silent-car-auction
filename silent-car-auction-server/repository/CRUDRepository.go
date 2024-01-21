// Package repository presents the functions that handle the database operations
package repository

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
    "silent-car-auction/model"
)

func toSQLString(value interface{}) string {
    switch value.(type) {
    case string:
        return fmt.Sprintf("'%v'", value)
    case int:
        return fmt.Sprintf("%v", value)
    case time.Time:
        date := value.(time.Time)
        return fmt.Sprintf("TO_DATE('%v', 'YYYY-MM-DD')", date.Format("2006-01-02"))
    }
    return ""
}


func appendWhere(sql string, conditions map[string]interface{}) string {
    if conditions == nil || len(conditions) == 0 {
        return sql
    }
    sql += " WHERE"
    for k, v := range conditions {
        sql += fmt.Sprintf(" %v=%v AND", k, toSQLString(v))
    }
    // remove last AND
    sql = sql[:len(sql) - 4]
    return sql
}

// SELECT * FROM table.Name() WHERE conditions;
func Select(table model.Table, conditions map[string]interface{}, c *gin.Context) {
    err := ensureDatabase()
    sql := fmt.Sprintf("SELECT * FROM %v", table.Name())
    sql = appendWhere(sql, conditions)
    sql += ";"
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

// UPDATE table.Name() SET column=value WHERE conditions;
func Update(table model.Table, column string, value interface{},
    conditions map[string]interface{}, c *gin.Context) {

    sql := fmt.Sprintf("UPDATE %v SET ", table.Name())
    sql += fmt.Sprintf("%v=%v", column, toSQLString(value))
    sql = appendWhere(sql, conditions)
    sql += ";"
    execSQL(sql, c)
}

// INSERT INTO table.Name() (table.Columns()) VALUES (values);
func Insert(table model.Table, values []interface{}, c *gin.Context) {
    sql := fmt.Sprintf("INSERT INTO %v (", table.Name())
    for _, column := range table.Columns() {
        sql += fmt.Sprintf("%v, ", column)
    }
    sql = sql[:len(sql) - 2] + ") VALUES ("
    for _, value := range values {
        sql += fmt.Sprintf("%v, ", toSQLString(value))
    }
    sql = sql[:len(sql) - 2] + ");"
    execSQL(sql, c)
}

// DELETE FROM table.Name() WHERE conditions;
func Delete(table model.Table, conditions map[string]interface{}, c *gin.Context) {
    sql := fmt.Sprintf("DELETE FROM %v", table.Name())
    sql = appendWhere(sql, conditions)
    sql += ";"
    execSQL(sql, c)
}

func execSQL(sql string, c *gin.Context) {
    fmt.Println(sql)
    err := ensureDatabase()
    if err != nil {
        c.String(500, err.Error())
        return
    }
    _, err = db.Exec(sql)
    if err != nil {
        c.String(400, err.Error())
        return
    }
    c.String(200, "success")
}
