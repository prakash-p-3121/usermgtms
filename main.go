package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/restlib"
	"github.com/prakash-p-3121/usermgtms/cfg"
	"github.com/prakash-p-3121/usermgtms/controller"
	"github.com/prakash-p-3121/usermgtms/database"
)

func main() {

	msConnectionsMap, err := restlib.CreateMsConnectionCfg("conf/microservice.toml")
	if err != nil {
		panic(err)
	}
	cfg.SetMsConnectionsMap(msConnectionsMap)

	connectionsMap, err := mysqllib.CreateShardConnectionsWithRetry(database.GetTableList())
	if err != nil {
		panic(err)
	}
	database.SetShardConnectionsMap(connectionsMap)

	db, err := mysqllib.CreateDatabaseConnectionWithRetryByCfg("conf/database.toml")
	if err != nil {
		panic(err)
	}
	database.SetSingleStoreConnection(db)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/usermgt/user", controller.UserCreate)

	err = router.Run("127.0.0.1:3000")
	if err != nil {
		panic("Error Starting UserMgtMs")
	}
}
