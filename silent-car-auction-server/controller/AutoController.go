package controller

import (
    "strconv"
    "time"
	"github.com/gin-gonic/gin"
    "silent-car-auction/model"
    "silent-car-auction/repository"
)

// Map of column names to types, used to convert query string values to the correct type
var autoTypeMap = map[string]string {
    "VIN": "string",
    "location": "string",
    "type_of_auto": "string",
    "mileage": "int",
    "year": "year",
}

// Function BindAutoController binds the auto controller to gin.Engine
func BindAutoController(router *gin.Engine) {
    router.GET("/auto", AutoGet)
    router.POST("/auto", AutoPostProcedure)
    router.PATCH("/auto/:column/*value", AutoPatch)
    router.DELETE("/auto", AutoDelete)
    router.GET("/auto/:brand", AutoGetBrand)
    router.GET("/auto/ford", AutoGetAllFords)
    router.GET("/auto/cheverolet", AutoGetAllCheverolets)
}

func AutoGet(c *gin.Context) {
    repository.Select(new(model.Auto), readConditions(autoTypeMap, c), c)
}

func AutoPost(c *gin.Context) {
    values := []interface{} {
        convToType(autoTypeMap, "VIN", c.PostForm("VIN")),
        convToType(autoTypeMap, "location", c.PostForm("location")),
        convToType(autoTypeMap, "type_of_auto", c.PostForm("type_of_auto")),
        convToType(autoTypeMap, "mileage", c.PostForm("mileage")),
        convToType(autoTypeMap, "year", c.PostForm("year")),
    }
    repository.Insert(new(model.Auto), values, c)
}

func AutoPatch(c *gin.Context) {
    column := c.Param("column")
    value := convToType(autoTypeMap, column, c.Param("value")[1:]) // remove leading slash
    repository.Update(new(model.Auto), column, value, readConditions(autoTypeMap, c), c)
}

func AutoDelete(c *gin.Context) {
    repository.Delete(new(model.Auto), readConditions(autoTypeMap, c), c)
}

func AutoGetBrand(c *gin.Context) {
    brand := c.Param("brand")
    repository.SelectByBrand(brand, c)
}

func AutoGetAllFords(c *gin.Context) {
    repository.SelectAllFords(c)
}

func AutoGetAllCheverolets(c *gin.Context) {
    repository.SelectAllCheverolets(c)
}

func AutoPostProcedure(c *gin.Context) {
    mileage, err := strconv.Atoi(c.PostForm("mileage"))
    if err != nil {
        c.String(400, err.Error())
        return
    }
    year, err := time.Parse("2006", c.PostForm("year"))
    if err != nil {
        c.String(400, err.Error())
        return
    }
    values := []interface{} {
        c.PostForm("VIN"),
        c.PostForm("location"),
        c.PostForm("type_of_auto"),
        mileage,
        year,
    }
    repository.AddAuto(values, c)
}