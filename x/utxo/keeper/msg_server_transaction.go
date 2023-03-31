package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	"utxo/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// トランザクションを受け取り、ビットコインにおけるトランザクション処理と同等の処理を行う
func (k msgServer) SendTransaction(goCtx context.Context, msg *types.MsgSendTransaction) (*types.MsgSendTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Make a hash of the transaction‘s data
	var transaction = msg.Tx
	data, _ := transaction.Marshal()
	txHash := sha256.Sum256(data)
	txHashStr := string(txHash[:])

	var inSum uint64 = 0

	// 各 Inputs について署名が正しいかを検証する
	for _, elm := range transaction.TxIn {
		// TxIn に対応する TxOut を検索する
		tx, found := k.GetTransaction(ctx, []byte(elm.PreviousOutpoint.Hash))
		if !found {
			return nil, fmt.Errorf("Not found transaction input utxo")
		}
		if len(tx.TxOut)+1 < int(elm.PreviousOutpoint.Index) {
			return nil, fmt.Errorf("Not exist transaction output index")
		}
		txOut := tx.TxOut[elm.PreviousOutpoint.Index]

		err := k.VerifyTxInput(elm, txOut)
		if err != nil {
			return nil, err
		}

		inSum += txOut.Value
	}

	// 各 Outputs の累計が inSum と等しいかを検証する
	var outSum uint64 = 0
	for _, elm := range transaction.TxOut {
		outSum += elm.Value
	}
	if inSum != outSum {
		return nil, fmt.Errorf("Invalid transaction")
	}

	// save tx to store
	k.AppendTransaction(
		ctx,
		txHashStr,
		*transaction,
	)

	return &types.MsgSendTransactionResponse{
		Hash: txHashStr,
	}, nil
}

// Verify transaction by checking the signature on bitcoin signature.
// If the signature is valid, the transaction is added to the utxo store.
func (k Keeper) VerifyTxInput(txIn *types.TxIn, txOut *types.TxOut) error {
	// Verify the signature by publicKey at secp256k1
	signature := txIn.SignatureScript
	publicKey := txOut.PkScript
	message, err := txIn.PreviousOutpoint.Marshal()
	if err != nil {
		return err
	}
	if !VerifySignature([]byte(signature), []byte(publicKey), message) {
		return fmt.Errorf("Invalid signature")
	}
	return nil
}

// verify signature by secp256k1
func VerifySignature(signature, publicKey, message []byte) bool {
	// Hash the message
	hash := sha256.Sum256(message)

	// Decode the provided public key
	var pubKey secp256k1.PubKey
	err := pubKey.Unmarshal(publicKey)
	if err != nil {
		return false
	}

	// Verify the signature using the decoded public key and the hashed message
	return pubKey.VerifySignature(hash[:], signature)
}
