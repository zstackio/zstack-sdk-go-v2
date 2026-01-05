// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectNfvInst operates on ReconnectNfvInst
func (cli *ZSClient) ReconnectNfvInst(uuid string, params param.ReconnectNfvInstParam) (*view.ReconnectNfvInstEventView, error) {
	resp := view.ReconnectNfvInstEventView{}
	if err := cli.Put("v1/vm-instances/appliances/nfvinst/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
