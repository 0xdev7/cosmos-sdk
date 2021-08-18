package middleware

import (
	"context"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type errorTxHandler struct {
	inner tx.TxHandler
	debug bool
}

// NewErrorTxMiddleware is a middleware that converts an error from inner
// middlewares into a abci.Response{Check,Deliver}Tx.
func NewErrorTxMiddleware(debug bool) tx.TxMiddleware {
	return func(txh tx.TxHandler) tx.TxHandler {
		return errorTxHandler{inner: txh, debug: debug}
	}
}

var _ tx.TxHandler = errorTxHandler{}

// CheckTx implements TxHandler.CheckTx.
func (txh errorTxHandler) CheckTx(ctx context.Context, tx sdk.Tx, req abci.RequestCheckTx) (abci.ResponseCheckTx, error) {
	res, err := txh.inner.CheckTx(ctx, tx, req)
	if err != nil {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		gInfo := sdk.GasInfo{GasUsed: sdkCtx.GasMeter().GasConsumed()}

		return sdkerrors.ResponseCheckTx(err, gInfo.GasWanted, gInfo.GasUsed, txh.debug), nil
	}

	return res, nil
}

// DeliverTx implements TxHandler.DeliverTx.
func (txh errorTxHandler) DeliverTx(ctx context.Context, tx sdk.Tx, req abci.RequestDeliverTx) (abci.ResponseDeliverTx, error) {
	res, err := txh.inner.DeliverTx(ctx, tx, req)
	if err != nil {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		gInfo := sdk.GasInfo{GasUsed: sdkCtx.GasMeter().GasConsumed()}

		return sdkerrors.ResponseDeliverTx(err, gInfo.GasWanted, gInfo.GasUsed, txh.debug), nil
	}

	return res, nil
}

// SimulateTx implements TxHandler.SimulateTx method.
func (txh errorTxHandler) SimulateTx(ctx context.Context, sdkTx sdk.Tx, req tx.RequestSimulateTx) (tx.ResponseSimulateTx, error) {
	res, err := txh.inner.SimulateTx(ctx, sdkTx, req)
	if err != nil {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		gInfo := sdk.GasInfo{GasUsed: sdkCtx.GasMeter().GasConsumed()}

		// In simulate mode, since the ResponseSimulateTx doesn't have
		// code/codespace/log, we return the error to baseapp.
		return tx.ResponseSimulateTx{GasInfo: gInfo, Result: res.Result}, err
	}

	return res, nil
}