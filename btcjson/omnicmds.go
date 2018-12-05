// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

type OmniGetInfoCmd struct {
}

func NewOmniGetInfoCmd() *OmniGetInfoCmd {
	return &OmniGetInfoCmd{}
}

type OmniGetBalanceCmd struct {
	Address    string
	Propertyid uint32
}

type OmniBalanceResult struct {
	Balance  string `json:"balance,omitempty"`
	Reserved string `json:"reserved,omitempty"`
	Frozen   string `json:"frozen,omitempty"`
}

func NewOmniGetBalanceCmd(addr string, propertyID uint32) *OmniGetBalanceCmd {
	return &OmniGetBalanceCmd{
		Address:    addr,
		Propertyid: propertyID,
	}
}

type OmniListBlockTxsCmd struct {
	Block uint32
}

func NewOmniListBlockTxsCmd(blockNum uint32) *OmniListBlockTxsCmd {
	return &OmniListBlockTxsCmd{
		Block: blockNum,
	}
}

type OmniGetTxCmd struct {
	TxID string
}

type OmniPropertyTx struct {
	PropertyId uint32
	Divisible  bool
	Amount     string
}

type OmniSendAllTx struct {
	Ecosystem string
	Subsends  []OmniPropertyTx
}

type OmniTxResult struct {
	Version          uint32
	Type_int         uint32
	Type             string
	TxID             string
	Fee              string
	Sendingaddress   string
	Referenceaddress string
	Ismine           bool
	Valid            bool
	Confirmations    uint32
	Block            uint32
	Positioninblock  uint32
	OmniPropertyTx
	OmniSendAllTx
}

func NewOmniGetTxCmd(txid string) *OmniGetTxCmd {
	return &OmniGetTxCmd{
		TxID: txid,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("omni_getinfo", (*OmniGetInfoCmd)(nil), flags)
	MustRegisterCmd("omni_getbalance", (*OmniGetBalanceCmd)(nil), flags)
	MustRegisterCmd("omni_listblocktransactions", (*OmniListBlockTxsCmd)(nil), flags)
	MustRegisterCmd("omni_gettransaction", (*OmniGetTxCmd)(nil), flags)
}
