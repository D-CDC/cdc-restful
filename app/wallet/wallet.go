package wallet

import (
	"cdc/backend/app/config"
	"crypto/ecdsa"
	"errors"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	client     *Client
	keydir     string //key store directory
	keyFile    string //the account key store file
	account    string
	privateKey *ecdsa.PrivateKey
}

var (
	mywallet *Wallet
)

func InitWallet() {
	mywallet = NewWallet(config.GetString("app.key_store_dir"))

	client := NewClient()
	client.Connect(config.GetString("app.ether_net_url"))

	if client.conn != nil {
		mywallet.client = client
	}
}

func NewWallet(keydir string) *Wallet {
	return &Wallet{
		keydir: keydir,
	}
}

func (w *Wallet) Create(password string) (string, error) {
	ks := keystore.NewKeyStore(w.keydir, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	w.account = account.Address.Hex()

	return w.account, nil
}

func (w *Wallet) importAccount(password string) (string, error) {
	var (
		key *keystore.Key
		err error
	)

	files, err := ioutil.ReadDir(w.keydir)
	if err != nil {
		log.Printf("Read %s directory err:%v\n", w.keydir, err)
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			filename := w.keydir + "/" + file.Name()
			key, err = Decrypt(filename, password)
			if err != nil {
				continue
			}

			//decrypt the private key successful
			w.privateKey = key.PrivateKey
			w.account = key.Address.Hex()
			w.keyFile = file.Name()
			return w.account, nil
		}
	}

	return "", err
}

func (w *Wallet) getPrivateKey(password string) (string, error) {
	filename := w.keydir + "/" + w.keyFile
	key, err := Decrypt(filename, password)
	if err != nil {
		return "", err
	}

	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)

	return hexutil.Encode(privateKeyBytes), nil
}

func (w *Wallet) getBalance() (string, error) {
	if w.client == nil {
		return "", errors.New("Please check network connection")
	}

	address := common.HexToAddress(w.account)
	balance, err := w.client.BalanceAt(address, nil)
	if err != nil {
		log.Println("Get balance err:", err)
		return "", err
	}

	return balance.String(), nil
}

func (w *Wallet) sendTransaction(toAccount string, amount, gasPrice *big.Int, gasLimit uint64) (string, error) {
	if w.client == nil {
		return "", errors.New("Please check network connection")
	}

	fromAddress := common.HexToAddress(w.account)

	nonce, err := w.client.PendingNonceAt(fromAddress)
	if err != nil {
		log.Println("Get nonce err:", err)
		return "", err
	}

	networkId, err := w.client.NetworkID()
	if err != nil {
		log.Println("Get network id err:", err)
		return "", err
	}

	var data []byte

	toAddress := common.HexToAddress(toAccount)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(networkId), w.privateKey)
	if err != nil {
		log.Println("Signed transaction err:", err)
		return "", err
	}

	err = w.client.SendTransaction(signedTx)
	if err != nil {
		log.Println("Send transaction err:", err)
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func (w *Wallet) getGasPrice() (string, error) {
	if w.client == nil {
		return "", errors.New("Please check network connection")
	}

	gasPrice, err := w.client.SuggestGasPrice()
	if err != nil {
		log.Println("Get suggest gas price err:", err)
		return "", err
	}

	return gasPrice.String(), nil
}
