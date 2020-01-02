package key

import (
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/CodeChain-io/codechain-sdk-go/vm"
)

func GetP2PKHBurnLockScript() []vm.OpCode {
	return []vm.OpCode{vm.COPY, 0x01, vm.BLAKE160, vm.EQ, vm.JZ, 0xff, vm.CHKSIG, vm.JZ, 0xff, vm.BURN}
}

func GetP2PKHBurnLockScriptHash() primitives.H160 {
	hash, _ := primitives.StringToH160("37572bdcc22d39a59c0d12d301f6271ba3fdd451")
	return hash
}

//TODO createAddress, createUnlockScript
