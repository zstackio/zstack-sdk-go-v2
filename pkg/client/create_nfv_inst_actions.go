// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateNfvInst creates NfvInst
func (cli *ZSClient) CreateNfvInst(params param.CreateNfvInstParam) (*view.CreateNfvInstEventView, error) {
	resp := view.CreateNfvInstEventView{}
	if err := cli.Post("v1/nfvinstgroup/inst", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
