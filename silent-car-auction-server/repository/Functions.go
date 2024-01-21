package repository

import (
	"github.com/gin-gonic/gin"
	"silent-car-auction/model/function"
)

func SelectByBrand(brand string, c *gin.Context) {
    callFunction(new(function.Brand), "select_by_brand", []interface{}{brand}, c)
}

func SelectAllFords(c *gin.Context) {
    callFunction(new(function.Brand), "avaliable_fords", nil, c)
}

func SelectAllCheverolets(c *gin.Context) {
    callFunction(new(function.Brand), "avaliable_cheverolets", nil, c)
}

func QueryMaxBids(c *gin.Context) {
    callFunction(new(function.MaxBid), "max_bid", nil, c)
}

func QueryWinnersAndLosersByAutoID(autoID int, c *gin.Context) {
    callFunction(new(function.WinnersAndLosers), "winners_and_losers", []interface{}{autoID}, c)
}
