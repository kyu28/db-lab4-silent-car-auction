// Package controller contains the functions that handle the requests to the server
package controller

import (
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
)

// Function readConditions reads the query string from the request and converts the values to the correct type
func readConditions(typeMap map[string]string, c *gin.Context) map[string]interface{} {
    conditions := make(map[string]interface{})
    for k, v := range c.Request.URL.Query() {
		fmt.Println(k, typeMap[k])
		if typeMap[k] == "" {
			continue
		}
		conditions[k] = convToType(typeMap, k, v[0])
    }
    return conditions
}

// Function convToType converts a key string to the correct type and value
func convToType(typeMap map[string]string, key string, str string) interface{} {
	var value interface{}
	var err error
	// No default fallthrough in Go
	switch typeMap[key] {
	case "int", "float":
		value, err = strconv.Atoi(str)
	case "string":
		value = str
	case "date":
		value, err = time.Parse("2006-01-02", str)
	case "year":
		value, err = time.Parse("2006", str)
	}
	if err != nil {
		return nil
	}
	return value
}