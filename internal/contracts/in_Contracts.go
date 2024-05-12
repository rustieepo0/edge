package contracts

import (
	"github.com/everFinance/goar"
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
