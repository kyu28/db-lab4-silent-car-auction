package controller

import (
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

// Map of column names to types, used to convert query string values to the correct type
var autoIndexTypeMap = map[string]string {
    "auto_id": "int",
    "VIN": "string",
}

// Function BindAutoIndexController binds the auto controller to gin.Engine
func BindAutoIndexController(router *gin.Engine) {
    router.GET("/auto-index", AutoIndexGet)
    router.POST("/auto-index", AutoIndexPost)
    router.PATCH("/auto-index/:column/*value", AutoIndexPatch)
    router.DELETE("/auto-index", AutoIndexDelete)
}

func AutoIndexGet(c *gin.Context) {
    repository.Select(new(model.AutoIndex), readConditions(autoIndexTypeMap, c), c)
}

func AutoIndexPost(c *gin.Context) {
    values := []interface{} {
        convToType(autoIndexTypeMap, "VIN", c.PostForm("VIN")),
    }
    repository.Insert(new(model.AutoIndex), values, c)
}

func AutoIndexPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(autoIndexTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.AutoIndex), column, value, readConditions(autoIndexTypeMap, c), c)
}

func AutoIndexDelete(c *gin.Context) {
    repository.Delete(new(model.AutoIndex), readConditions(autoIndexTypeMap, c), c)
}