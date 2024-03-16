package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/restlib"
	"github.com/prakash-p-3121/usermgtms/cfg"
	"github.com/prakash-p-3121/usermgtms/controller"
	"github.com/prakash-p-3121/usermgtms/database"
	"log"
)

func main() {

	msConnectionsMap, err := restlib.CreateMsConnectionCfg("conf/microservice.toml")
	if err != nil {
		panic(err)
	}
	cfg.SetMsConnectionsMap(msConnectionsMap)

	cfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		panic(err)
	}
	connectionsMap, err := mysqllib.CreateShardConnectionsWithRetry(database.GetShardedTableList(), cfg.Host, cfg.Port)
	if err != nil {
		panic(err)
	}
	log.Println(connectionsMap)
	database.SetShardConnectionsMap(connectionsMap)

	db, err := mysqllib.CreateDatabaseConnectionWithRetryByCfg("conf/database.toml")
	if err != nil {
		panic(err)
	}
	database.SetSingleStoreConnection(db)

	router := gin.Default()
	routerGroup := router.Group("/usermgt")
	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup.POST("/v1/user", controller.CreateUser)
	routerGroup.GET("/v1/user", controller.FindUser)

	err = router.Run("127.0.0.1:3004")
	if err != nil {
		panic("Error Starting UserMgtMs")
	}
}
