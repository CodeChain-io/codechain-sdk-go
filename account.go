package rpc

type Account struct {
	rpcClient rpcClient
}

func (acc *Account) GetList() []string {
	const method = "account_getList"
	var accounts []string
	acc.rpcClient.call(callInterface{method: method, id: ""}, &accounts)
	return accounts
}

func (acc *Account) Create(passphrase string) string {
	const method = "account_create"
	var platformAddress string
	acc.rpcClient.call(callInterface{method: method}, &platformAddress, passphrase)
	return platformAddress
}

func (acc *Account) ImportRaw(secret string, passphrase string) string {
	const method = "account_importRaw"
	var platformAddress string
	acc.rpcClient.call(callInterface{method: method}, &platformAddress, secret, passphrase)
	return platformAddress

}

func (acc *Account) Unlock(account string, passphrase string, duration int) {
	const method = "account_unlock"
	acc.rpcClient.call(callInterface{method: method}, nil, account, passphrase, duration)
}

func (acc *Account) Sign(message string, account string, passphrase string) string {
	const method = "account_sign"
	var signature string
	acc.rpcClient.call(callInterface{method: method}, &signature, message, account, passphrase)
	return signature
}

func (acc *Account) SendTransaction(transaction UnsignedTransaction, account string, passphrase string) interface{} {
	const method = "account_sendTransaction"
	var result interface{}
	acc.rpcClient.call(callInterface{method: method}, &result, transaction, account, passphrase)
	return result
}

func (acc *Account) ChangePassword(account string, oldPassphrase string, newPassphrase string) {
	const method = "account_changePassword"
	acc.rpcClient.call(callInterface{method: method}, nil, account, oldPassphrase, newPassphrase)
}
