package key

import (
	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/CodeChain-io/codechain-sdk-go/vm"
)

func GetP2PKHLockScript() []vm.OpCode {
	return []vm.OpCode{vm.COPY, 0x01, vm.BLAKE160, vm.EQ, vm.JZ, 0xff, vm.CHKSIG}
}

func GetP2PKHLockScriptHash() primitives.H160 {
	hash, _ := primitives.StringToH160("5f5960a7bca6ceeeb0c97bc717562914e7a1de04")
	return hash
}

func CreateP2PKHUnlockScript(pubkey []byte, privkey []byte, txHash primitives.H256) []byte {

	encodedTag := []byte{3}

	signature := crypto.SignEcdsa(txHash.Bytes(), privkey)
	pushb := byte(vm.PUSHB)

	result := []byte{pushb, 65}

	result = append(result, signature...)
	result = append(result, pushb)
	result = append(result, byte(len(encodedTag)))
	result = append(result, encodedTag...)
	result = append(result, pushb)
	result = append(result, 64)
	result = append(result, pubkey...)

	return result
}

// TODO createAddress
