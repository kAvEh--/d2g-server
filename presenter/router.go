package presenter

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	gin.SetMode(gin.DebugMode)
	r.MaxMultipartMemory = 50 << 20 // 50 MB

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//Version 1 APIs
	// APIs without authentication
	v1 := r.Group("/api/v1")
	{
		v1.GET("/login")
	}

	return r
}
