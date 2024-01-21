package repository

import (
	"github.com/gin-gonic/gin"
)

func AddCustomer(values []interface{}, c *gin.Context) {
    callProcedure("add_customer", values, c)
}

func AddAuto(values []interface{}, c *gin.Context) {
    callProcedure("add_auto", values, c)
}

func AddTransaction(values []interface{}, c *gin.Context) {
    callProcedure("add_transaction", values, c)
}
