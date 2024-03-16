package main

import (
	"fmt"
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
	connectionsMap, err := mysqllib.CreateShardConnectionsWithRetry(database.GetTableList(), cfg.Host, cfg.Port)
	if err != nil {
		panic(err)
	}
	log.Println(connectionsMap)
	database.SetShardConnectionsMap(connectionsMap)

	connectionsMap.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})

	//db, err := mysqllib.CreateDatabaseConnectionWithRetryByCfg("conf/database.toml")
	//if err != nil {
	//	panic(err)
	//}
	//database.SetSingleStoreConnection(db)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/usermgt/user", controller.CreateUser)
	router.GET("/usermgt/user", controller.FindUser)

	err = router.Run("127.0.0.1:3000")
	if err != nil {
		panic("Error Starting UserMgtMs")
	}
}
