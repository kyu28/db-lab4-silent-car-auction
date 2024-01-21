package controller

import (
    "strconv"
    "time"
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

var transactionTypeMap = map[string]string {
    "transaction_id": "int",
    "auto_id": "int",
    "customer_id": "int",
    "bid": "float",
    "bid_date": "date",
}

func BindTransactionController(router *gin.Engine) {
    router.GET("/transaction", TransactionGet)
    router.POST("/transaction", TransactionPostProcedure)
    router.PATCH("/transaction/:column/*value", TransactionPatch)
    router.DELETE("/transaction", TransactionDelete)
    router.GET("/transaction/max-bid", TransactionGetMaxBid)
}

func TransactionGet(c *gin.Context) {
    repository.Select(new(model.Transaction), readConditions(transactionTypeMap, c), c)
}

func TransactionPost(c *gin.Context) {
    values := []interface{} {
        convToType(transactionTypeMap, "transaction_id", c.PostForm("transaction_id")),
    }
    repository.Insert(new(model.Transaction), values, c)
}

func TransactionPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(transactionTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.Transaction), column, value, readConditions(transactionTypeMap, c), c)
}

func TransactionDelete(c *gin.Context) {
    repository.Delete(new(model.Transaction), readConditions(transactionTypeMap, c), c)
}

func TransactionGetMaxBid(c *gin.Context) {
    repository.QueryMaxBids(c)
}

func TransactionPostProcedure(c *gin.Context) {
    keys := []string {
        "auto_id",
        "customer_id",
        "bid",
    }
    values := make([]interface{}, 4)
    var err error
    for i := 0; i < 3; i++ {
        values[i], err = strconv.Atoi(c.PostForm(keys[i]))
        if err != nil {
            c.String(400, err.Error())
            return
        }
    }
    values[3], err = time.Parse("2006-01-02", c.PostForm("bid_date"))
    if err != nil {
        c.String(400, err.Error())
        return
    }
    repository.AddTransaction(values, c)
}