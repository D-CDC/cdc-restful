package api

import (
	"cdc/backend/app/wallet"
	"log"

	"github.com/gin-gonic/gin"
)

func getPassword(c *gin.Context) string {
	param := make(map[string]string)

	err := c.Bind(&param)
	if err != nil {
		return ""
	}

	return param["password"]
}

func Login(c *gin.Context) {
	password := getPassword(c)
	if password == "" {
		InternalError(c, -1, "Param parse error")
		return
	}

	log.Println("the passowrd:", password)

	account, err := wallet.ImportAccount(password)
	if err != nil {
		InternalError(c, -1, "Incorrect password")
		return
	}

	data := make(map[string]string)
	data["account"] = account
	SuccessResponse(c, 0, "Login success", data)
}

func CreateAccount(c *gin.Context) {
	password := getPassword(c)
	if password == "" {
		InternalError(c, -1, "Param parse error")
		return
	}

	log.Println("the passowrd:", password)
	account, err := wallet.CreateAccount(password)
	if err != nil {
		InternalError(c, -1, "Create a account failed")
		return
	}

	data := make(map[string]string)
	data["account"] = account
	SuccessResponse(c, 0, "Create success", data)
}

func DerivePrivateKey(c *gin.Context) {
	password := getPassword(c)
	if password == "" {
		InternalError(c, -1, "Param parse error")
		return
	}

	privateKey, err := wallet.DerivePrivateKey(password)
	if err != nil {
		InternalError(c, -1, "Incorrect password")
		return
	}

	data := make(map[string]string)
	data["private_key"] = privateKey
	SuccessResponse(c, 0, "Create success", data)
}

func GetBalance(c *gin.Context) {
	balance, err := wallet.GetBalance()
	if err != nil {
		InternalError(c, -1, err.Error())
		return
	}

	data := make(map[string]string)
	data["balance"] = balance
	SuccessResponse(c, 0, "Get balance success", data)
}

func GetPrint(c *gin.Context) {
	data := make(map[string]string)
	data["balance"] = "hello"
	SuccessResponse(c, 0, "Get balance success", data)
}
