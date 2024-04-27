package Config

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type AssetManagementAsset struct {
	Id           *big.Int
	Name         string
	Description  string
	OwnerAddress common.Address
	IsApproved   bool
}
