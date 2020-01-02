package vm

type OpCode byte

const (
	NOP OpCode = iota
	BURN
	SUCCESS
	FAIL
)

const (
	NOT OpCode = iota + 0x10
	EQ
)

const (
	JMP OpCode = iota + 0x20
	JNZ
	JZ
)

const (
	PUSH OpCode = iota + 0x30
	POP
	PUSHB
	DUP
	SWAP
	COPY
	DROP
)

const (
	CHKSIG OpCode = iota + 0x80
	CHKMULTISIG
)

const (
	BLAKE256 OpCode = iota + 0x90
	SHA256
	RIPEMD160
	KECCAK256
	BLAKE160
)

const (
	BLKNUM OpCode = iota + 0xa0
)

const (
	CHKTIMELOCK OpCode = iota + 0xb0
)

var opCodeToString = map[OpCode]string{
	NOP:         "NOP",
	BURN:        "BURN",
	SUCCESS:     "SUCCESS",
	FAIL:        "FAIL",
	NOT:         "NOT",
	EQ:          "EQ",
	JMP:         "JMP",
	JNZ:         "JNZ",
	JZ:          "JZ",
	PUSH:        "PUSH",
	PUSHB:       "PUSHB",
	DUP:         "DUP",
	SWAP:        "SWAP",
	COPY:        "COPY",
	DROP:        "DROP",
	CHKSIG:      "CHKSIG",
	CHKMULTISIG: "CHKMULTISIG",
	BLAKE256:    "BLAKE256",
	SHA256:      "SHA256",
	RIPEMD160:   "RIPEMD160",
	KECCAK256:   "KECCAK256",
	BLAKE160:    "BLAKE160",
	BLKNUM:      "BLKNUM",
	CHKTIMELOCK: "CHKTIMELOCK",
}

func (op OpCode) String() string {
	str := opCodeToString[op]
	return str
}

var stringToOp = map[string]OpCode{
	"NOP":         NOP,
	"BURN":        BURN,
	"SUCCESS":     SUCCESS,
	"FAIL":        FAIL,
	"NOT":         NOT,
	"EQ":          EQ,
	"JMP":         JMP,
	"JNZ":         JNZ,
	"JZ":          JZ,
	"PUSH":        PUSH,
	"PUSHB":       PUSHB,
	"DUP":         DUP,
	"SWAP":        SWAP,
	"COPY":        COPY,
	"DROP":        DROP,
	"CHKSIG":      CHKSIG,
	"CHKMULTISIG": CHKMULTISIG,
	"BLAKE256":    BLAKE256,
	"SHA256":      SHA256,
	"RIPEMD160":   RIPEMD160,
	"KECCAK256":   KECCAK256,
	"BLAKE160":    BLAKE160,
	"BLKNUM":      BLKNUM,
	"CHKTIMELOCK": CHKTIMELOCK,
}

func StringToOp(str string) OpCode {
	return stringToOp[str]
}
