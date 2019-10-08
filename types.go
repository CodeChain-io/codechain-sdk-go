package rpc

type UnsignedTransaction struct {
	fee       int
	networkId string
	seq       int
	action    interface{}
}

type Transaction struct {
	blockHash        string
	blockNumber      int
	transactionIndex int
	result           bool
	seq              int
	fee              int
	hash             string
	networkId        string
	sig              string
	action           string
} // TODO Handle action details

type Block struct {
	Author           string        `json:"author"`
	ExtraData        []int         `json:"extraData"`
	Hash             string        `json:"hash"`
	Number           int           `json:"number"`
	ParentHash       string        `json:"parentHash"`
	Score            string        `json:"score"`
	Seal             [][]int       `json:"seal"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        int           `json:"timestamp"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
}

type Text struct {
	content   string
	certifier string
}
