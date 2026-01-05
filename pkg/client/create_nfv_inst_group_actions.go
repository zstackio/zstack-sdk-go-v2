// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateNfvInstGroup creates NfvInstGroup
func (cli *ZSClient) CreateNfvInstGroup(params param.CreateNfvInstGroupParam) (*view.CreateNfvInstGroupEventView, error) {
	resp := view.CreateNfvInstGroupEventView{}
	if err := cli.Post("v1/nfvinstgroup/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
