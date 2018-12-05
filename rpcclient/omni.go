// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
)

func (c *Client) OMNIGetInfo() (string, error) {
	cmd := btcjson.NewOmniGetInfoCmd()
	res, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (c *Client) OMNIGetBalance(addr string, propertyID uint32) (*btcjson.OmniBalanceResult, error) {
	cmd := btcjson.NewOmniGetBalanceCmd(addr, propertyID)
	res, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return nil, err
	}

	var result btcjson.OmniBalanceResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) OMNIListBlockTxs(blockNum uint32) ([]string, error) {
	cmd := btcjson.NewOmniListBlockTxsCmd(blockNum)
	res, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) OMNIGetTx(txid string) (*btcjson.OmniTxResult, error) {
	cmd := btcjson.NewOmniGetTxCmd(txid)
	res, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return nil, err
	}

	var result btcjson.OmniTxResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
