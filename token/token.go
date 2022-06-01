package token

import (
	"github.com/jxlwqq/chainbase-client-go/api"
	"time"
)

type Client struct {
	apiClient *api.Client
}

func New(apiClient *api.Client) *Client {
	return &Client{apiClient: apiClient}
}

func (c *Client) GetBalance() {
	_ = "account/balance"
}

type Metadata struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	TotalSupply     string `json:"total_supply"`
}

type AccountToken struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	Balance         string `json:"balance"`
}

type AccountTransaction struct {
	Type                 int       `json:"type"`
	Status               int       `json:"status"`
	BlockNumber          int       `json:"block_number"`
	BlockTimestamp       time.Time `json:"block_timestamp"`
	TransactionHash      string    `json:"transaction_hash"`
	TransactionIndex     int       `json:"transaction_index"`
	FromAddress          string    `json:"from_address"`
	ToAddress            string    `json:"to_address"`
	Value                string    `json:"value"`
	Input                string    `json:"input"`
	Nonce                int       `json:"nonce"`
	ContractAddress      string    `json:"contract_address"`
	Gas                  int       `json:"gas"`
	GasPrice             int64     `json:"gas_price"`
	GasUsed              int       `json:"gas_used"`
	EffectiveGasPrice    int64     `json:"effective_gas_price"`
	CumulativeGasUsed    int       `json:"cumulative_gas_used"`
	MaxFeePerGas         int       `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas int       `json:"max_priority_fee_per_gas"`
	ChainId              int       `json:"chain_id"`
}

type Price struct {
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Client) GetTokenMetadata() {
	_ = "token/metadata"
}

func (c *Client) GetAccountTokens() {

	_ = "account/tokens"
}

func (c *Client) GetAccountTransactions() {

	_ = "account/txs"
}

func (c *Client) GetTokenHolders() {

	_ = "token/holders"
}

func (c *Client) GetTokenPrice() {

	_ = "token/price"
}

func (c *Client) GetTokenPriceHistory() {

	_ = "token/price/history"
}
