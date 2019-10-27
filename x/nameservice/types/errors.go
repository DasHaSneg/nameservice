package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace     sdk.CodespaceType = ModuleName
	CodeNameDoesNitExist sdk.CodeType      = 101
)

func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNitExist, "Name does not exist")
}
