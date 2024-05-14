package contracts

import (
	"encoding/json"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/liteseed/aogo"
)

type Context struct {
	ao      *aogo.AO
	process string
	signer  *goar.ItemSigner
}

func NewContext(ao *aogo.AO, process string, signer *goar.ItemSigner) *Context {
	return &Context{
		ao:      ao,
		process: process,
		signer:  signer,
	}
}

type Contract struct {
	*Context
}

func NewContract(ao *aogo.AO, process string, signer *goar.ItemSigner) *Contract {
	return &Contract{
		Context: NewContext(ao, process, signer),
	}
}

type GetUploadResponse struct {
	Status   string `json:"status"`
	Quantity string `json:"quantity"`
	Block    string `json:"block"`
	Bundler  string `json:"bundler"`
}

func (c *Contract) GetUpload(dataItemId string) (*GetUploadResponse, error) {
	mID, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Upload"}, {Name: "DataItemId", Value: dataItemId}}, "", c.signer)
	if err != nil {
		return nil, err
	}

	result, err := c.ao.ReadResult(c.process, mID)
	if err != nil {
		return nil, err
	}

	var response GetUploadResponse
	err = json.Unmarshal([]byte(result.Messages[0]["Data"].(string)), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Contract) Notify(dataItemId string, transactionID string) error {
	_, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Notify"}, {Name: "DataItemId", Value: dataItemId}, {Name: "TransactionId", Value: transactionID}}, "", c.signer)
	return err
}

func (c *Contract) Release(dataItemId string) error {
	_, err := c.ao.SendMessage(c.process, "release", []types.Tag{{Name: "Action", Value: "Release"}, {Name: "DataItemId", Value: dataItemId}}, "", c.signer)
	return err
}

func (c *Contract) Stake(url string) error {
	_, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Stake"}, {Name: "Url", Value: url}}, "", c.signer)
	return err
}

func (c *Contract) GetStaker() (string, error) {
	mID, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Staked"}}, "", c.signer)
	if err != nil {
		return "", err
	}

	result, err := c.ao.ReadResult(c.process, mID)
	if err != nil {
		return "", err
	}

	return result.Messages[0]["Data"].(string), nil
}

func (c *Contract) Unstake() error {
	_, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Unstake"}}, "", c.signer)
	return err
}

type GetBalanceResponse struct {
	Balance string `json:"balance"`
}

func (c *Contract) GetBalance() (string, error) {
	mID, err := c.ao.SendMessage(c.process, "", []types.Tag{{Name: "Action", Value: "Balance"}}, "", c.signer)
	if err != nil {
		return "", err
	}

	result, err := c.ao.ReadResult(c.process, mID)
	if err != nil {
		return "", err
	}

	var response GetBalanceResponse
	err = json.Unmarshal([]byte(result.Messages[0]["Data"].(string)), &response)
	if err != nil {
		return "", err
	}

	return response.Balance, nil
}
