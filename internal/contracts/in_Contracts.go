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
