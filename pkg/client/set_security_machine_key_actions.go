// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetSecurityMachineKey operates on SetSecurityMachineKey
func (cli *ZSClient) SetSecurityMachineKey(params param.SetSecurityMachineKeyParam) (*view.SetSecurityMachineKeyEventView, error) {
	resp := view.SetSecurityMachineKeyEventView{}
	if err := cli.Post("v1/secret-resource-pool-token/set/{uuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
