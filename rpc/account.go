package rpc

type Account struct {
	rpcClient rpcClient
}

func (acc *Account) GetList() ([]string, error) {
	const method = "account_getList"
	var accounts []string
	err := acc.rpcClient.call(callInterface{method: method, id: ""}, &accounts)
	return accounts, err
}

func (acc *Account) Create(passphrase string) (string, error) {
	const method = "account_create"
	var platformAddress string
	err := acc.rpcClient.call(callInterface{method: method}, &platformAddress, passphrase)
	return platformAddress, err
}

func (acc *Account) ImportRaw(secret string, passphrase string) (string, error) {
	const method = "account_importRaw"
	var platformAddress string
	err := acc.rpcClient.call(callInterface{method: method}, &platformAddress, secret, passphrase)
	return platformAddress, err
}

func (acc *Account) Unlock(account string, passphrase string, duration int) error {
	const method = "account_unlock"
	err := acc.rpcClient.call(callInterface{method: method}, nil, account, passphrase, duration)
	return err
}

func (acc *Account) Sign(message string, account string, passphrase string) (string, error) {
	const method = "account_sign"
	var signature string
	err := acc.rpcClient.call(callInterface{method: method}, &signature, message, account, passphrase)
	return signature, err
}

func (acc *Account) SendTransaction(transaction UnsignedTransaction, account string, passphrase string) (interface{}, error) {
	const method = "account_sendTransaction"
	var result interface{}
	err := acc.rpcClient.call(callInterface{method: method}, &result, transaction, account, passphrase)
	return result, err
}

func (acc *Account) ChangePassword(account string, oldPassphrase string, newPassphrase string) error {
	const method = "account_changePassword"
	err := acc.rpcClient.call(callInterface{method: method}, nil, account, oldPassphrase, newPassphrase)
	return err
}
