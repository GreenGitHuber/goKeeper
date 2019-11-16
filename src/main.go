package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// ExampleNewClient1()
	// ExampleNewClient2()
	
	// demo 
	// r := gin.Default()
    // r.GET("/ping", func(c *gin.Context) {
    //     c.JSON(200, gin.H{
    //         "message": "pong",
    //     })
    // })
	// r.Run() // listen and serve on 0.0.0.0:8080

	// 无参数
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Who are you?")
	// })
	// r.Run()

	router := gin.Default()
	// /heartbeat?service=echo&expire=1000
	router.GET("/heartbeat", func(c *gin.Context) {
		service_name := c.DefaultQuery("service", "Guest")
		ip := c.Query("ip")
		port := c.Query("port")
		expire_time := c.Query("expire") 
		c.String(http.StatusOK, "service name %s, ip %s, port %s, expire %s", service_name, ip, port, expire_time)
	})

	// /getService?name=echo
	router.GET("/getService", func(c *gin.Context) {
		service_name := c.Query("name")

		// get data 
		name := ExampleNewClient3()

		c.String(http.StatusOK, "service name %s, get name %s", service_name, name)
	})
	
	router.Run(":8080")
}
