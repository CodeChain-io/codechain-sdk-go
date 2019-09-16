package rpc

type UnsingnedTransaction struct {
	fee       int
	networkId string
	seq       int
	action    interface{}
}

type Account struct {
	rpcClient rpcClient
}

func (acc *Account) GetList() []string {
	const method = "account_getList"
	response := acc.rpcClient.call(callInterface{method: method, id: ""})
	return response.Result.([]string)
}

func (acc *Account) Create(passphrase string) interface{} {
	const method = "account_create"
	response := acc.rpcClient.call(callInterface{method: method}, passphrase)
	return response.Result
}

func (acc *Account) ImportRaw(secret string, passphrase string) string {
	const method = "account_importRaw"
	response := acc.rpcClient.call(callInterface{method: method}, secret, passphrase)
	return response.Result.(string)
}

func (acc *Account) Unlock(account string, passphrase string, duration int) {
	const method = "account_unlock"
	acc.rpcClient.call(callInterface{method: method}, account, passphrase, duration)
}

func (acc *Account) sign(message string, account string, passphrase string) string {
	const method = "account_sign"
	response := acc.rpcClient.call(callInterface{method: method}, message, account, passphrase)
	return response.Result.(string)
}

func (acc *Account) sendTransaction(transaction UnsingnedTransaction, account string, passphrase string) interface{} {
	const method = "account_sendTransaction"
	response := acc.rpcClient.call(callInterface{method: method}, transaction, account, passphrase)
	return response.Result
}

func (acc *Account) changePassword(account string, oldPassphrase string, newPassphrase string) {
	const method = "account_changePassword"
	acc.rpcClient.call(callInterface{method: method}, account, oldPassphrase, newPassphrase)
}
