package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"nameservice/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, k)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, k)
		case QueryNames:
			return queryNames(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	value := k.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(k.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	whois := k.GetWhois(ctx, path[0])

	res, err := codec.MarshalJSONIndent(k.cdc, whois)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var namesList types.QueryResNames

	iterator := k.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(k.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
