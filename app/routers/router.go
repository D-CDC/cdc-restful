package routers

import (
	"cdc/backend/app/routers/api"
	"cdc/backend/app/routers/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Throttle(1000, 20))

	r.POST("/api/login", api.Login)

	account := r.Group("/api/account")
	{
		account.POST("/create", api.CreateAccount)
		account.POST("/private_key", api.DerivePrivateKey)
		account.GET("/balance", api.GetBalance)
		account.GET("/print", api.GetPrint)

	}

	transaction := r.Group("/api/transaction")
	{
		transaction.POST("/send", api.SendTransaction)
		transaction.POST("/trx", api.SendTrx)
		transaction.GET("/gas_price", api.GetGasPrice)
		transaction.GET("/trxHash", api.GetTrxInfo)
	}

	system := r.Group("/api/system")
	{
		system.GET("/node", api.GetNodeInfo)
		system.GET("/blockNumber", api.GetBlockInfo)
		system.GET("/threads", api.StartMine)
		system.GET("/miner", api.StopMine)
	}

	return r
}
