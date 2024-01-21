package controller

import (
    "strconv"
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

var customerTypeMap = map[string]string {
    "first_name": "string",
    "last_name": "string",
    "location": "string",
    "email": "string",
}

func BindCustomerController(router *gin.Engine) {
    router.GET("/customer", CustomerGet)
    router.POST("/customer", CustomerPostProcedure)
    router.PATCH("/customer/:column/*value", CustomerPatch)
    router.DELETE("/customer", CustomerDelete)
    router.GET("/customer/winners-and-losers/:auto_id", CustomerGetWinnersAndLosers)
}

func CustomerGet(c *gin.Context) {
    repository.Select(new(model.Customer), readConditions(customerTypeMap, c), c)
}

func CustomerPost(c *gin.Context) {
    values := []interface{} {
        convToType(customerTypeMap, "first_name", c.PostForm("first_name")),
        convToType(customerTypeMap, "last_name", c.PostForm("last_name")),
        convToType(customerTypeMap, "location", c.PostForm("location")),
        convToType(customerTypeMap, "email", c.PostForm("email")),
    }
    repository.Insert(new(model.Customer), values, c)
}

func CustomerPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(customerTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.Customer), column, value, readConditions(customerTypeMap, c), c)
}

func CustomerDelete(c *gin.Context) {
    repository.Delete(new(model.Customer), readConditions(customerTypeMap, c), c)
}

func CustomerGetWinnersAndLosers(c *gin.Context) {
    autoID, err := strconv.Atoi(c.Param("auto_id"))
    if err != nil {
        c.String(400, err.Error())
        return
    }
    repository.QueryWinnersAndLosersByAutoID(autoID, c)
}

func CustomerPostProcedure(c *gin.Context) {
    values := []interface{} {
        c.PostForm("first_name"),
        c.PostForm("last_name"),
        c.PostForm("location"),
        c.PostForm("email"),
    }
    repository.AddCustomer(values, c)
}