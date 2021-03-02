
package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"

	"gitlab.com/kenzchiro/bitkub_challange/go/src/solution"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	res:=solution.DateFormatter("29/02/2020")
	r.GET("/ping", func(c *gin.Context) {
	 c.JSON(200, gin.H{
	  "message": res,
	 })
	})
	r.Run() // listen and serve on 0.0.0.0:8000
   }