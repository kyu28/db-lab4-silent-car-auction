package controller

import (
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

var winnerTypeMap = map[string]string {
    "auto_id": "int",
    "customer_id": "int",
}

func BindWinnerController(router *gin.Engine) {
    router.GET("/winner", WinnerGet)
    router.POST("/winner", WinnerPost)
    router.PATCH("/winner/:column/*value", WinnerPatch)
    router.DELETE("/winner", WinnerDelete)
}

func WinnerGet(c *gin.Context) {
    repository.Select(new(model.Winner), readConditions(winnerTypeMap, c), c)
}

func WinnerPost(c *gin.Context) {
    values := []interface{} {
        convToType(winnerTypeMap, "auto_id", c.PostForm("auto_id")),
        convToType(winnerTypeMap, "customer_id", c.PostForm("customer_id")),
    }
    repository.Insert(new(model.Winner), values, c)
}

func WinnerPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(winnerTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.Winner), column, value, readConditions(winnerTypeMap, c), c)
}

func WinnerDelete(c *gin.Context) {
    repository.Delete(new(model.Winner), readConditions(winnerTypeMap, c), c)
}