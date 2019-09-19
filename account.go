package rpc

type UnsingnedTransaction struct {
	fee       int
	networkId string
	seq       int
	action    interface{}
}

func GetList() []string {
	const method = "account_getList"
	response := call(callInterface{method: method, id: ""})
	return response.Result.([]string)
}

func Create(passphrase string) interface{} {
	const method = "account_create"
	response := call(callInterface{method: method}, passphrase)
	return response.Result
}

func ImportRaw(secret string, passphrase string) string {
	const method = "account_importRaw"
	response := call(callInterface{method: method}, secret, passphrase)
	return response.Result.(string)
}

func Unlock(account string, passphrase string, duration int) {
	const method = "account_unlock"
	call(callInterface{method: method}, account, passphrase, duration)
}

func sign(message string, account string, passphrase string) string {
	const method = "account_sign"
	response := call(callInterface{method: method}, message, account, passphrase)
	return response.Result.(string)
}

func sendTransaction(transaction UnsingnedTransaction, account string, passphrase string) interface{} {
	const method = "account_sendTransaction"
	response := call(callInterface{method: method}, transaction, account, passphrase)
	return response.Result
}

func changePassword(account string, oldPassphrase string, newPassphrase string) {
	const method = "account_changePassword"
	call(callInterface{method: method}, account, oldPassphrase, newPassphrase)
}
