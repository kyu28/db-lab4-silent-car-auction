package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "fmt"
    "os"
    "silent-car-auction/controller"
)

func main() {
    fmt.Println("--- SILENT CAR AUCTION SERVER ---")
    port := ":8000"
    if len(os.Args) > 1 {
        port = ":" + os.Args[1]
    } else {
        fmt.Println("No port specified, using default port 8000")
    }
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PATCH", "GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
    controller.BindAutoController(router)
    controller.BindAutoIndexController(router)
    controller.BindCustomerController(router)
    controller.BindCustomerIndexController(router)
    controller.BindTransactionController(router)
    controller.BindWinnerController(router)
    router.Run(port)
}