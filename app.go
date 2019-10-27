package app

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types/module"
	"os"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/cosmos/cosmos-sdk/x/genaccounts"

	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tm-db"

	"nameservice/x/nameservice"
)

const appName = "nameservice"

var (
	// default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.nscli")

	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.nsd")

	// ModuleBasicManager is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager()
)

type nameServiceApp struct {
	*bam.BaseApp
	cdc *codec.Codec
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameServiceApp{
		bApp,
		cdc,
	}

	return app
}
