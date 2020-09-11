package wallet

func CreateAccount(password string) (string, error) {
	return mywallet.Create(password)
}

func ImportAccount(password string) (string, error) {
	return mywallet.importAccount(password)
}

func DerivePrivateKey(password string) (string, error) {
	return mywallet.getPrivateKey(password)
}

func GetBalance() (string, error) {
	return mywallet.getBalance()
}
