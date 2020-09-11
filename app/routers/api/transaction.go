package api

import (
	"cdc/backend/app/wallet"

	"github.com/gin-gonic/gin"
)

func SendTransaction(c *gin.Context) {
	param := make(map[string]string)
	err := c.Bind(&param)

	transHash, err := wallet.SendTransaction(param)
	if err != nil {
		InternalError(c, -1, err.Error())
		return
	}

	data := make(map[string]string)
	data["transaction_hash"] = transHash

	SuccessResponse(c, 0, "Send transaction success", data)
}

func GetGasPrice(c *gin.Context) {
	gasPrice, err := wallet.GetGasPrice()
	if err != nil {
		InternalError(c, -1, err.Error())
		return
	}

	data := make(map[string]string)
	data["gas_price"] = gasPrice
	SuccessResponse(c, 0, "Get gas price success", data)
}
