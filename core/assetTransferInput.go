package core

import (
	"errors"
)

type Timelock struct {
	TimelockType string `json:"type"`
	Value        int    `json:"value"`
}

type AssetTransferInputJSON struct {
	PrevOut      AssetOutPointJSON `json:"prevOut"`
	Timelock     Timelock          `json:"timelock"`
	LockScript   []int             `json:"lockScript"`
	UnlockScript []int             `json:"unlockScript"`
}

type AssetTransferInput struct {
	PrevOut      AssetOutPoint
	Timelock     *Timelock
	LockScript   []byte
	UnlockScript []byte
}

func NewAssetTransferInput(prevOut AssetOutPoint, timelock *Timelock, lockScript []byte, unlockScript []byte) AssetTransferInput {
	return AssetTransferInput{prevOut, timelock, lockScript, unlockScript}
}

func (a AssetTransferInput) ToEncodeObject() []interface{} {
	encodedTimelock, _ := ConvertTimelockToEncodeObject(a.Timelock)
	return []interface{}{
		a.PrevOut.ToEncodeObject(),
		encodedTimelock,
		a.LockScript,
		a.UnlockScript}
}

func (a AssetTransferInput) ToJSON() AssetTransferInputJSON {
	lockScript := make([]int, len(a.LockScript))
	for i, d := range a.LockScript {
		res := int(d)
		lockScript[i] = res
	}

	unlockScript := make([]int, len(a.UnlockScript))
	for i, d := range a.UnlockScript {
		res := int(d)
		unlockScript[i] = res
	}
	return AssetTransferInputJSON{a.PrevOut.ToJSON(), *a.Timelock, lockScript, unlockScript}
}

func (a AssetTransferInput) WithOutScript() AssetTransferInput {
	return AssetTransferInput{a.PrevOut, a.Timelock, []byte{}, []byte{}}
}

func (a *AssetTransferInput) SetLockScript(lockScript []byte) {
	a.LockScript = lockScript
}

func (a *AssetTransferInput) SetUnlockScript(unlockScript []byte) {
	a.UnlockScript = unlockScript
}

func ConvertTimelockToEncodeObject(t *Timelock) ([]interface{}, error) {
	if t == nil {
		return []interface{}{}, nil
	}

	var typeEncoded int
	switch timelockType := t.TimelockType; timelockType {
	case "block":
		typeEncoded = 1
	case "blockAge":
		typeEncoded = 2
	case "time":
		typeEncoded = 3
	case "timeAge":
		typeEncoded = 4
	default:
		return nil, errors.New("Unexpected timelock type: " + timelockType)
	}
	return []interface{}{[]interface{}{typeEncoded, t.Value}}, nil
}
