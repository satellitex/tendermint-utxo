package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgSendTransaction = "send_transaction"
)

var _ sdk.Msg = &MsgSendTransaction{}

func NewMsgSendTransaction(tx *Transaction) *MsgSendTransaction {
	return &MsgSendTransaction{
		Tx: tx,
	}
}

func (msg *MsgSendTransaction) Route() string {
	return RouterKey
}

func (msg *MsgSendTransaction) Type() string {
	return TypeMsgSendTransaction
}

func (msg *MsgSendTransaction) GetSigners() []sdk.AccAddress {
	res := []sdk.AccAddress{}
	for _, elm := range msg.Tx.TxOut {
		signer, err := sdk.AccAddressFromBech32(elm.PkScript)
		if err != nil {
			panic(err)
		}
		res = append(res, signer)
	}
	return res
}

func (msg *MsgSendTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendTransaction) ValidateBasic() error {
	// _, err := sdk.AccAddressFromBech32(msg.Tx.TxOut.PkScript)
	// if err != nil {
	// 	return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	return nil
}
