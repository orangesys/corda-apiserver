package api

import (
	"github.com/gin-gonic/gin"
	"orangesys.io/corda-apiserver/internal/app/corda"
)

//Init ...
func Init() {
	cordaSrv := corda.New()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	{
		corda := v1.Group("/corda")
		{
			corda.GET("/nodeconf", cordaSrv.GetNodeConf)
			corda.GET("/certs", cordaSrv.GetCerts)
		}
	}

	r.Run(":8080")
}
