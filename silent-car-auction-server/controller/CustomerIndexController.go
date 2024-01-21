package controller

import (
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

var customerIndexTypeMap = map[string]string {
    "customer_id": "int",
    "email": "string",
}

func BindCustomerIndexController(router *gin.Engine) {
    router.GET("/customer-index", CustomerIndexGet)
    router.POST("/customer-index", CustomerIndexPost)
    router.PATCH("/customer-index/:column/*value", CustomerIndexPatch)
    router.DELETE("/customer-index", CustomerIndexDelete)
}

func CustomerIndexGet(c *gin.Context) {
    repository.Select(new(model.CustomerIndex), readConditions(customerIndexTypeMap, c), c)
}

func CustomerIndexPost(c *gin.Context) {
    values := []interface{} {
        convToType(customerIndexTypeMap, "email", c.PostForm("email")),
    }
    repository.Insert(new(model.CustomerIndex), values, c)
}

func CustomerIndexPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(customerIndexTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.CustomerIndex), column, value, readConditions(customerIndexTypeMap, c), c)
}

func CustomerIndexDelete(c *gin.Context) {
    repository.Delete(new(model.CustomerIndex), readConditions(customerIndexTypeMap, c), c)
}