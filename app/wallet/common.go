package wallet

import (
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Decrypt(filename, password string) (*keystore.Key, error) {
	storeData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Read %s file err:%v\n", filename, err)
		return nil, err
	}

	key, err := keystore.DecryptKey(storeData, password)
	if err != nil {
		log.Printf("Keystore decrypt key err:%v\n", err)
		return nil, err
	}

	return key, nil
}

func StrToBigInt(value string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(value, 16)
	if !ok {
		log.Println("Big int SetString error")
		return nil
	}

	return n
}

func StrOToBigInt(value string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		log.Println("Big int SetString error")
		return nil
	}

	return n
}
