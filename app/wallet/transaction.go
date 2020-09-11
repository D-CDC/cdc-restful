package wallet

import (
	"errors"
	"strconv"
)

func SendTransaction(param map[string]string) (string, error) {
	toAccount := param["to_account"]
	amount := StrToBigInt(param["amount"])
	if amount == nil {
		return "", errors.New("amount param parse error")
	}

	gasPrice := StrToBigInt(param["gas_price"])
	if amount == nil {
		return "", errors.New("gas_price param parse error")
	}

	gasLimit, err := strconv.Atoi(param["gas_limit"])
	if err != nil {
		return "", errors.New("gas_limit param parse error")
	}

	return mywallet.sendTransaction(toAccount, amount, gasPrice, uint64(gasLimit))
}

func GetGasPrice() (string, error) {
	return mywallet.getGasPrice()
}
