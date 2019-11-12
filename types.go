package rpc

type UnsignedTransaction struct {
	Fee       int    `json:"fee"`
	NetworkID string `json:"network_id"`
	Seq       int    `json:"seq"`
	Action    Action `json:"action"`
}

type AssetScheme interface{}
type Asset interface{}
type CommonParams interface{}

type Text struct {
	Content   string `json:"content"`
	Certifier string `json:"certifier"`
}

type Action interface {
	getType() string
}

type MintAsset struct {
	Type                string      `json:"type"`
	NetworkID           string      `json:"network_id"`
	ShardID             int         `json:"shard_id"`
	Metadata            string      `json:"metadata"`
	Approver            string      `json:"appprover,omitempty"`
	Registrar           string      `json:"registrar,omitempty"`
	AllowedScriptHashes []string    `json:"allowed_script_hashes"`
	Output              interface{} `json:"output"`
	Approvals           []string    `json:"approvals"`
	Tracker             string      `json:"tracker"`
}

type TransferAsset struct {
	Type       string        `json:"type"`
	NetworkID  string        `json:"network_id"`
	Burns      []interface{} `json:"burns"`
	Inputs     []interface{} `json:"inputs"`
	Outputs    []interface{} `json:"outputs"`
	Orders     []interface{} `json:"orders"`
	Metadata   string        `json:"metadata"`
	Approvals  []string      `json:"approvals"`
	Tracker    string        `json:"tracker"`
	Expiration int           `json:"expiration,omitempty"`
}

type ChangeAssetScheme struct {
	Type string `json:"type"`
}

type IncreaseAssetSupply struct {
	Type string `json:"type"`
}

type UnwrapCCC struct {
	Type string `json:"type"`
}

type WrapCCC struct {
	Type string `json:"type"`
}

type Pay struct {
	Type string `json:"type"`
}

type SetRegularKey struct {
	Type string `json:"type"`
}

type Store struct {
	Type string `json:"type"`
}

type Remove struct {
	Type string `json:"type"`
}

type Custom struct {
	Type      string `json:"type"`
	NetworkID string `json:"network_id"`
	HandlerID string `json:"handler_id"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      int    `json:"blockNumber"`
	TransactionIndex int    `json:"transactionIndex"`
	Result           bool   `json:"result"`
	Seq              int    `json:"seq"`
	Fee              string `json:"fee"`
	Hash             string `json:"hash"`
	NetworkID        string `json:"networkID"`
	Sig              string `json:"sig"`
	Action           Action `json:"action"`
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

func (t MintAsset) getType() string {
	return "mintAsset"
}

func (t TransferAsset) getType() string {
	return "transferAsset"
}

func (t ChangeAssetScheme) getType() string {
	return "changeAssetScheme"
}

func (t IncreaseAssetSupply) getType() string {
	return "increaseAssetSupply"
}

func (t UnwrapCCC) getType() string {
	return "unwrapCCC"
}

func (t WrapCCC) getType() string {
	return "wrapCCC"
}

func (t Pay) getType() string {
	return "pay"
}

func (t SetRegularKey) getType() string {
	return "setRegularKey"
}

func (t Store) getType() string {
	return "store"
}

func (t Remove) getType() string {
	return "remove"
}

func (t Custom) getType() string {
	return "custom"
}
